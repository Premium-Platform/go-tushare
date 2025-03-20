package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// Level 日志级别
type Level int

const (
	// DEBUG 调试级别
	DEBUG Level = iota
	// INFO 信息级别
	INFO
	// WARN 警告级别
	WARN
	// ERROR 错误级别
	ERROR
	// FATAL 致命级别
	FATAL
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

// Logger 日志记录器
type Logger struct {
	level  Level
	writer io.Writer
	logger *log.Logger
	mu     sync.Mutex
}

var (
	// 默认日志记录器
	defaultLogger = NewLogger(os.Stdout, INFO)
)

// NewLogger 创建一个新的日志记录器
func NewLogger(writer io.Writer, level Level) *Logger {
	return &Logger{
		level:  level,
		writer: writer,
		logger: log.New(writer, "", 0),
	}
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// GetLevel 获取日志级别
func (l *Logger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// SetOutput 设置日志输出
func (l *Logger) SetOutput(writer io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writer = writer
	l.logger.SetOutput(writer)
}

// formatLog 格式化日志消息
func (l *Logger) formatLog(level Level, format string, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	msg := fmt.Sprintf(format, args...)
	return fmt.Sprintf("[%s] [%s] %s", timestamp, levelNames[level], msg)
}

// log 记录日志
func (l *Logger) log(level Level, format string, args ...interface{}) {
	if level < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	formattedMsg := l.formatLog(level, format, args...)
	l.logger.Println(formattedMsg)

	// 如果是致命错误，终止程序
	if level == FATAL {
		os.Exit(1)
	}
}

// Debug 调试级别日志
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info 信息级别日志
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn 警告级别日志
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error 错误级别日志
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Fatal 致命级别日志，会导致程序终止
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
}

// 全局函数
// SetLevel 设置默认日志级别
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// SetOutput 设置默认日志输出
func SetOutput(writer io.Writer) {
	defaultLogger.SetOutput(writer)
}

// Debug 全局调试级别日志
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info 全局信息级别日志
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warn 全局警告级别日志
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Error 全局错误级别日志
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// Fatal 全局致命级别日志
func Fatal(format string, args ...interface{}) {
	defaultLogger.Fatal(format, args...)
}

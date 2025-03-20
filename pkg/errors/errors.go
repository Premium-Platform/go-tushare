package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	// ErrInvalidToken 表示无效的令牌
	ErrInvalidToken = errors.New("invalid token")

	// ErrNetworkFailure 表示网络请求失败
	ErrNetworkFailure = errors.New("network request failed")

	// ErrAPILimit 表示API调用超出限制
	ErrAPILimit = errors.New("api call limit exceeded")

	// ErrInvalidParameter 表示参数无效
	ErrInvalidParameter = errors.New("invalid parameter")

	// ErrServerError 表示服务器内部错误
	ErrServerError = errors.New("server internal error")

	// ErrUnknown 表示未知错误
	ErrUnknown = errors.New("unknown error")
)

// APIError 表示API响应的错误
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error 实现error接口
func (e *APIError) Error() string {
	return fmt.Sprintf("API error (code: %d): %s", e.Code, e.Message)
}

// NewAPIError 创建一个新的API错误
func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

// Wrap 包装一个错误，增加上下文信息
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Wrapf 包装一个错误，增加格式化的上下文信息
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// Cause 获取错误链中的原始错误
func Cause(err error) error {
	return errors.Cause(err)
}

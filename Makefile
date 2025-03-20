.PHONY: build clean test run fmt

# 默认目标
all: fmt build

# 构建示例
build:
	go build -o bin/stock_example cmd/stock_example/main.go
	go build -o bin/basic_example cmd/basic_example/main.go
	go build -o bin/bar_example cmd/bar_example/main.go

# 运行示例
run: build
	./bin/example

# 运行测试
test:
	go test -v ./...

# 代码格式化
fmt:
	go fmt ./...

# 清理编译产物
clean:
	rm -rf bin/
	go clean

# 安装依赖
deps:
	go mod tidy
	go mod download 
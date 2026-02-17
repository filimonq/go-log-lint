BINARY_NAME=gologlint
PLUGIN_NAME=gologlint.so
CMD_PATH=cmd/go-log-lint/main.go
PLUGIN_PATH=plugin/main.go

.PHONY: all build build-plugin test clean run fix help

all: test build build-plugin

build:
	@echo "🔨 Building CLI tool..."
	go build -o $(BINARY_NAME) $(CMD_PATH)

build-plugin:
	@echo "🔌 Building Plugin..."
	go build -buildmode=plugin -o $(PLUGIN_NAME) $(PLUGIN_PATH)

test:
	@echo "🧪 Running tests..."
	go test -v ./...

run: build
	@echo "🚀 Running linter..."
	./$(BINARY_NAME) ./...

fix: build
	@echo "🔧 Running linter with FIX..."
	./$(BINARY_NAME) -fix ./...

clean:
	@echo "🧹 Cleaning up..."
	rm -f $(BINARY_NAME) $(PLUGIN_NAME)
	go clean

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
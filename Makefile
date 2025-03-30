.PHONY: build clean

# Default target
all: build

# Create bin directory and build the binary
build:
	mkdir -p bin
	go build -o bin/mcp-nutanix

# Clean build artifacts
clean:
	rm -rf bin/

install:
	go install golang.org/x/tools/cmd/goimports@latest

# Generate resource and tool implementations
generate:
	# Copy the main.go file to cmd/codegen if it doesn't exist
	go run internal/codegen/cmd/main.go --output .
	goimports -w ./pkg/resources
	goimports -w ./pkg/tools
	go mod tidy
	@echo "Generation complete!"

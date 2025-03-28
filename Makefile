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

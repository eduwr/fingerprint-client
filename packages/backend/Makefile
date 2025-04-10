.PHONY: build run test clean lint format

# Variables
BINARY_NAME=backend
GO=go
GOFMT=gofmt
GOLINT=golangci-lint
BUILD_DIR=build
VERSION=$(shell git describe --tags --always --dirty)

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)"

# Default target
all: build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) main.go

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	$(GO) run main.go

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v -race -cover ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	$(GO) clean

# Run linter
lint:
	@echo "Running linter..."
	$(GOLINT) run

# Format code
format:
	@echo "Formatting code..."
	$(GOFMT) -w .

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy
	$(GO) mod download

# Install development tools
dev-tools:
	@echo "Installing development tools..."
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run in development mode with hot reload
dev:
	@echo "Starting development server with hot reload..."
	air

# Generate swagger documentation
swagger:
	@echo "Generating swagger documentation..."
	swag init -g main.go 
#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored messages
print_message() {
    echo -e "${2}${1}${NC}"
}

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check for required tools
check_tools() {
    local tools=("go" "make" "air")
    local missing_tools=()

    for tool in "${tools[@]}"; do
        if ! command_exists "$tool"; then
            missing_tools+=("$tool")
        fi
    done

    if [ ${#missing_tools[@]} -ne 0 ]; then
        print_message "Missing required tools: ${missing_tools[*]}" "$RED"
        print_message "Please install them before continuing." "$YELLOW"
        exit 1
    fi
}

# Install development tools
install_dev_tools() {
    print_message "Installing development tools..." "$YELLOW"
    make dev-tools
    go install github.com/cosmtrek/air@latest
}

# Setup development environment
setup() {
    print_message "Setting up development environment..." "$YELLOW"
    check_tools
    install_dev_tools
    make deps
}

# Run the application in development mode
run_dev() {
    print_message "Starting development server..." "$GREEN"
    make dev
}

# Run tests
run_tests() {
    print_message "Running tests..." "$YELLOW"
    make test
}

# Run linter
run_linter() {
    print_message "Running linter..." "$YELLOW"
    make lint
}

# Format code
format_code() {
    print_message "Formatting code..." "$YELLOW"
    make format
}

# Main script
case "$1" in
    "setup")
        setup
        ;;
    "dev")
        run_dev
        ;;
    "test")
        run_tests
        ;;
    "lint")
        run_linter
        ;;
    "format")
        format_code
        ;;
    *)
        print_message "Usage: $0 {setup|dev|test|lint|format}" "$YELLOW"
        exit 1
        ;;
esac 
# Define variables
PROJECT_NAME := case-study/library
BUILD_DIR := ./bin
SRC_DIR := ./cmd/$(PROJECT_NAME)
MAIN_FILE := ./app/main.go

# Default target
.PHONY: all
all: build

# Build the project
.PHONY: build
build: 
	@echo "Building the project..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(PROJECT_NAME) $(MAIN_FILE)

# Run tests
.PHONY: test
test: 
	@echo "Running tests..."
	go test ./...

# Clean up build files
.PHONY: clean
clean: 
	@echo "Cleaning up..."
	go clean
	rm -rf $(BUILD_DIR)

# Run the application
.PHONY: run
run: 
	@echo "Running the application..."
	go run $(MAIN_FILE)

# Display help
.PHONY: help
help: 
	@echo "Makefile commands:"
	@echo "  all          - Build the project"
	@echo "  build        - Build the binary"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean up build files"
	@echo "  run          - Run the application"
	@echo "  help         - Display this help message"
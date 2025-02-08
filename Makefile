# Project Variables
APP_NAME=crud
BUILD_DIR=bin
GO_FILES=$(shell find . -name '*.go')
JSON_STORAGE_MODE="json"
MONGO_STORAGE_MODE="mongo"

# Default storage mode (JSON)
export STORAGE_MODE=$(JSON_STORAGE_MODE)

# Docker Variables
MONGO_CONTAINER=mongodb
MONGO_IMAGE=mongo:latest
MONGO_PORT=27017

# Default target (runs the app in JSON mode)
.DEFAULT_GOAL := run-json

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy

# Build the Go application (with main.go inside cmd/)
build:
	@echo "Building $(APP_NAME)..."
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/main.go

# Run the application in JSON mode (default)
run-json: export STORAGE_MODE=$(JSON_STORAGE_MODE)
run-json: build
	@echo "Running $(APP_NAME) in JSON mode..."
	@$(BUILD_DIR)/$(APP_NAME)

# Run the application in MongoDB mode
run-mongo: export STORAGE_MODE=$(MONGO_STORAGE_MODE)
run-mongo: build start-mongo
	@echo "Running $(APP_NAME) in MongoDB mode..."
	@$(BUILD_DIR)/$(APP_NAME)

# Start MongoDB container using Docker
start-mongo:
	@echo "Starting MongoDB in Docker..."
	docker run -d --name $(MONGO_CONTAINER) -p $(MONGO_PORT):27017 $(MONGO_IMAGE)

# Stop and remove MongoDB container
stop-mongo:
	@echo "Stopping MongoDB..."
	docker stop $(MONGO_CONTAINER) && docker rm $(MONGO_CONTAINER)

# Run with Docker Compose (MongoDB mode)
run-docker:
	@echo "Running with Docker Compose (MongoDB mode)..."
	docker compose up --build

# Clean up build files
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Help command
help:
	@echo "Available commands:"
	@echo "  make install       - Install dependencies"
	@echo "  make build         - Build the Go app"
	@echo "  make run-json      - Run in JSON mode (default)"
	@echo "  make run-mongo     - Run in MongoDB mode"
	@echo "  make start-mongo   - Start MongoDB in Docker"
	@echo "  make stop-mongo    - Stop MongoDB container"
	@echo "  make run-docker    - Run using Docker Compose"
	@echo "  make clean         - Remove build files"

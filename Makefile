# Makefile for Go CRUD App

# Variables
APP_NAME = go-crud-app
GO_CMD = go
DOCKER_COMPOSE_CMD = docker-compose

# Install dependencies
.PHONY: init
init:
	$(GO_CMD) mod init crud
	$(GO_CMD) mod tidy

# Build the Go application
.PHONY: build
build:
	$(GO_CMD) build -o $(APP_NAME) ./cmd/main.go

# Run the application locally (without Docker)
.PHONY: run
run:
	$(GO_CMD) run cmd/main.go

# Run tests
.PHONY: test
test:
	$(GO_CMD) test ./...

# Run with Docker Compose
.PHONY: up
docker-up:
	$(DOCKER_COMPOSE_CMD) up --build -d

# Stop Docker services
.PHONY: down
docker-down:
	$(DOCKER_COMPOSE_CMD) down

# Clean up build artifacts
.PHONY: clean
clean:
	rm -f $(APP_NAME)

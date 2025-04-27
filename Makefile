# Makefile for go-clean-architecture project

# Variables
GO := go
DOCKER_COMPOSE := docker-compose
SWAG := $(GOPATH)/bin/swag

# Default target
.PHONY: all
all: build

# Install dependencies
.PHONY: deps
deps:
	$(GO) mod download
	$(GO) install github.com/swaggo/swag/cmd/swag@latest
	$(GO) install github.com/air-verse/air@latest

# Clean generated Swagger docs
.PHONY: clean-swagger
clean-swagger:
	rm -rf docs/docs.go docs/swagger.json docs/swagger.yaml

# Generate Swagger documentation
.PHONY: swagger
swagger: clean-swagger
	$(SWAG) init

# Build the Docker image
.PHONY: build
build: swagger
	$(DOCKER_COMPOSE) build

# Start the application with Docker Compose
.PHONY: up
up: build
	$(DOCKER_COMPOSE) up -d

# Stop and remove Docker containers
.PHONY: down
down:
	$(DOCKER_COMPOSE) down

# Clean Docker resources (containers, images, volumes) and temporary files
.PHONY: clean
clean: clean-swagger
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans
	rm -rf tmp

# Run the application locally without Docker (using Air)
.PHONY: run-local
run-local: swagger
	air -c .air.toml

# Run tests (assuming tests exist in the project)
.PHONY: test
test:
	$(GO) test ./... -v

# Apply database migrations (if using a migration tool like goose or migrate)
.PHONY: migrate
migrate:
	@echo "Applying database migrations..."
	# Add your migration command here, e.g., using goose or migrate
	# Example with goose:
	# goose -dir infrastructure/database/migrations postgres "host=localhost port=5432 user=EngenMe password=Mehmed_0793727673 dbname=GO_CLEAN_ARCHITECTURE sslmode=disable" up

# View logs for Docker containers
.PHONY: logs
logs:
	$(DOCKER_COMPOSE) logs -f

# Help command to display available targets
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make deps          Install dependencies"
	@echo "  make clean-swagger Clean generated Swagger documentation"
	@echo "  make swagger       Generate Swagger documentation"
	@echo "  make build         Build Docker images"
	@echo "  make up            Start application with Docker Compose"
	@echo "  make down          Stop and remove Docker containers"
	@echo "  make clean         Clean Docker resources and temporary files"
	@echo "  make run-local     Run application locally with Air"
	@echo "  make test          Run tests"
	@echo "  make migrate       Apply database migrations"
	@echo "  make logs          View Docker container logs"
	@echo "  make help          Display this help message"
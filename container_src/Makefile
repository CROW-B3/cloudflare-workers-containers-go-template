.PHONY: help run build test clean dev install-air air migrate docker-build docker-run fmt lint

# Default target
help:
	@echo "Available commands:"
	@echo "  make run           - Run the application"
	@echo "  make build         - Build the application"
	@echo "  make dev           - Run with hot reload (Air)"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make install-air   - Install Air for hot reload"
	@echo "  make migrate       - Run database migrations"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter"

# Run the application
run:
	@echo "Starting application..."
	go run main.go

# Build the application
build:
	@echo "Building application..."
	go build -o bin/app main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf tmp/
	rm -f coverage.out

# Install Air for hot reload
install-air:
	@echo "Installing Air..."
	go install github.com/air-verse/air@latest

# Run with hot reload using Air
dev: install-air
	@echo "Starting development server with hot reload..."
	air

# Alternative: Run air directly (assumes air is already installed)
air:
	@echo "Starting Air..."
	air

# Database migration (uncomment when ready)
migrate:
	@echo "Running database migrations..."
	go run main.go migrate

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Install linter
install-lint:
	@echo "Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t go-microservice:latest .

# Docker run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env go-microservice:latest

# Docker compose up
docker-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

# Docker compose down
docker-down:
	@echo "Stopping services..."
	docker-compose down

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download

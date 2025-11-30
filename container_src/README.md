# Go Microservice with Gin

A production-ready Go microservice built with industry-standard packages and best practices.

## Tech Stack

- **Framework**: Gin (most popular Go web framework in 2025)
- **Configuration**: Viper (with YAML and environment variable support)
- **Logging**: Zap (high-performance structured logging)
- **Validation**: go-playground/validator
- **Database**: GORM with PostgreSQL
- **Metrics**: Prometheus
- **Hot Reload**: Air
- **Containerization**: Docker & Docker Compose

## Project Structure

```
.
├── cmd/                    # Application entrypoints
├── config/                 # Configuration files
│   └── config.yaml
├── internal/               # Private application code
│   ├── handlers/          # HTTP handlers
│   ├── middleware/        # Custom middleware
│   ├── models/            # Data models
│   ├── repositories/      # Data access layer
│   ├── routes/            # Route definitions
│   └── services/          # Business logic
├── pkg/                    # Public libraries
│   ├── database/          # Database connection
│   ├── logger/            # Logger setup
│   ├── response/          # Standard responses
│   └── validator/         # Validation helpers
├── .air.toml              # Air configuration
├── .env.example           # Environment variables example
├── docker-compose.yaml    # Docker Compose setup
├── Dockerfile             # Container image
├── Makefile               # Build commands
└── main.go                # Application entry point
```

## Features

- Clean architecture with separation of concerns
- Request ID tracking for distributed tracing
- Structured logging with Zap
- Prometheus metrics for monitoring
- Health check endpoints (/health, /ready, /live)
- CORS middleware
- Panic recovery middleware
- Input validation
- Database connection pooling
- Graceful shutdown
- Hot reload for development

## Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL (optional, for database features)
- Docker & Docker Compose (optional)

### Installation

1. Clone the repository
2. Copy environment file:
   ```bash
   cp .env.example .env
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

### Running the Application

#### Development (with hot reload):
```bash
make dev
```

#### Without hot reload:
```bash
make run
```

#### Using Docker Compose:
```bash
make docker-up
```

This will start:
- Go application on port 8080
- PostgreSQL on port 5432
- Prometheus on port 9090

## Available Endpoints

### Health Checks
- `GET /health` - Basic health check
- `GET /ready` - Readiness check (includes DB connectivity)
- `GET /live` - Liveness check
- `GET /metrics` - Prometheus metrics

### API v1
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - List users (with pagination)
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

## Configuration

Configuration can be set via:
1. `config/config.yaml` file
2. Environment variables (will override config file)
3. `.env` file (loaded automatically)

Environment variables use uppercase with underscores:
- `SERVER_PORT` overrides `server.port`
- `DATABASE_HOST` overrides `database.host`

## Make Commands

```bash
make help           # Show available commands
make run            # Run the application
make dev            # Run with hot reload (Air)
make build          # Build binary
make test           # Run tests
make test-coverage  # Run tests with coverage
make clean          # Clean build artifacts
make fmt            # Format code
make lint           # Run linter
make docker-build   # Build Docker image
make docker-run     # Run Docker container
make docker-up      # Start with Docker Compose
make docker-down    # Stop Docker Compose services
```

## Database Migrations

To enable database migrations, uncomment the migration code in `main.go`:

```go
if err := database.DB.AutoMigrate(&models.User{}); err != nil {
    logger.Fatal("Failed to migrate database")
}
```

## Monitoring

Prometheus metrics are exposed at `/metrics`. The metrics include:
- HTTP request count (by method, endpoint, status)
- HTTP request duration (histogram)
- Custom business metrics (add your own)

Access Prometheus UI at http://localhost:9090 when using Docker Compose.

## Production Deployment

1. Set environment to production:
   ```bash
   export APP_ENVIRONMENT=production
   export SERVER_MODE=release
   ```

2. Build optimized binary:
   ```bash
   make build
   ```

3. Or use Docker:
   ```bash
   make docker-build
   make docker-run
   ```

## Best Practices Implemented

- Structured logging with context
- Request ID for tracing
- Graceful shutdown
- Configuration management
- Database connection pooling
- Input validation
- Error handling
- Health checks
- Metrics collection
- Clean architecture
- Dependency injection

## License

MIT

# Go Clean Architecture Backend

A production-ready backend API built with Go, following Clean Architecture principles, CQRS pattern, and a modular design. The project provides user management and authentication functionality, with a focus on scalability, maintainability, and best practices.

## Technologies Used

- **Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/) with PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **API Documentation**: [Swaggo](https://github.com/swaggo/swag) for Swagger/OpenAPI
- **Hot Reloading**: [Air](https://github.com/air-verse/air)
- **Containerization**: Docker and Docker Compose
- **Build Automation**: Make (via `Makefile`)

## Features

- **Clean Architecture**: Organized into layers (`api`, `application`, `domain`, `infrastructure`, `interfaces`) for clear separation of concerns.
- **User Management**: Full CRUD operations for users, including email lookup.
- **Authentication**: JWT-based authentication for securing protected endpoints.
- **API Documentation**: Auto-generated Swagger UI for easy API exploration.
- **Database Initialization**: Automatic database creation and schema setup via migrations.
- **Best Practices**: Structured error handling, logging middleware, and environment-based configuration.

## Project Structure

```
.
├── api
│   ├── handlers
│   │   ├── auth_handler.go          # Authentication endpoint handlers
│   │   └── user_handler.go          # User endpoint handlers
│   ├── middlewares
│   │   ├── auth_middleware.go       # JWT authentication middleware
│   │   └── logging_middleware.go    # Request logging middleware
│   └── routes
│       └── routes.go                # API route definitions
├── application
│   ├── commands
│   │   ├── create_user.go           # Command for creating users
│   │   ├── delete_user.go           # Command for deleting users
│   │   └── update_user.go           # Command for updating users
│   ├── queries
│   │   ├── get_user_by_email.go     # Query for fetching user by email
│   │   ├── get_user_by_id.go        # Query for fetching user by ID
│   │   └── get_users.go             # Query for fetching all users
│   └── services
│       ├── auth_service.go          # Authentication business logic
│       └── user_service.go          # User management business logic
├── docker-compose.yml               # Docker Compose configuration
├── Dockerfile                       # Docker build configuration
├── docs
│   ├── docs.go                      # Generated Swagger documentation
│   ├── swagger.json                 # Swagger JSON spec
│   └── swagger.yaml                 # Swagger YAML spec
├── domain
│   └── entities
│       └── user.go                  # User entity definition
├── go.mod                           # Go module dependencies
├── go.sum                           # Go module checksums
├── infrastructure
│   ├── database
│   │   ├── connection.go            # Database connection setup
│   │   ├── generic_repository.go    # Generic repository implementation
│   │   ├── migrations
│   │   │   └── 001_create_users_table.sql  # Database schema migration
│   │   └── postgres_user_repository.go     # User-specific repository
│   └── utils
│       ├── config.go                # Environment configuration
│       ├── errors.go                # Custom error handling
│       └── jwt.go                   # JWT utility functions
├── interfaces
│   └── repositories
│       ├── generic_repository.go    # Generic repository interface
│       └── user_repository.go       # User repository interface
├── main.go                          # Application entry point
├── Makefile                         # Build and run automation
├── README.md                        # Project documentation
└── tmp
    ├── build-errors.log             # Build error logs
    └── main                         # Compiled binary (via Air)
```

## Prerequisites

- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) (for Docker setup)
- [Go](https://golang.org/dl/) 1.24 or higher (for local development)
- [Make](https://www.gnu.org/software/make/) (optional, for using `Makefile`)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-clean-architecture.git
cd go-clean-architecture
```

### Environment Setup

1. **Copy the example environment file**:
   The repository includes a `.env` file with default settings. If you need to customize (e.g., change `JWT_SECRET`), create a new `.env`:
   ```bash
   cp .env .env.local
   ```
   Edit `.env.local` as needed (e.g., update `DB_PASSWORD` or `JWT_SECRET`).

   Default `.env`:
   ```env
   # Server settings
    PORT=8080
    ENV=development
    
    # Database settings
    DB_HOST=postgres
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=password
    DB_NAME=cleanarchdb
    DB_SSL_MODE=disable
    
    # JWT settings
    JWT_SECRET=your_jwt_secret_key
    JWT_EXPIRATION_HOURS=24
   ```

2. **Ensure `.env` is not committed**:
   The `.gitignore` file already excludes `.env` to prevent committing sensitive data.

### Running the Application

#### With Docker (Recommended)

1. **Build and start the application**:
   ```bash
   docker-compose up --build
   ```
   Or, if using the `Makefile`:
   ```bash
   make up
   ```

   This will:
    - Build the Go application Docker image.
    - Create and initialize the `GO_CLEAN_ARCHITECTURE` PostgreSQL database with the `users` table (via `infrastructure/database/migrations/001_create_users_table.sql`).
    - Start the API server on `http://localhost:8080`.

2. **Access the Swagger UI**:
   Open your browser and navigate to:
   ```
   http://localhost:8080/api/v1/swagger/index.html
   ```

3. **Stop the application**:
   ```bash
   docker-compose down
   ```
   Or:
   ```bash
   make down
   ```
   To reset the database (clear all data):
   ```bash
   docker-compose down -v
   ```

#### Without Docker

1. **Install dependencies**:
   ```bash
   make deps
   ```
   Or:
   ```bash
   go mod download
   go install github.com/swaggo/swag/cmd/swag@v1.16.3
   go install github.com/air-verse/air@latest
   ```

   **Note**: Install `swag@v1.16.3` to avoid Swagger compatibility issues with `LeftDelim` and `RightDelim`.

2. **Set up PostgreSQL locally**:
    - Install PostgreSQL (e.g., via `brew install postgresql` on macOS or `sudo apt-get install postgresql` on Ubuntu).
    - Create the database:
      ```bash
      psql -U postgres -c "CREATE DATABASE GO_CLEAN_ARCHITECTURE;"
      ```
    - Apply the migration:
      ```bash
      psql -U EngenMe -d GO_CLEAN_ARCHITECTURE -f infrastructure/database/migrations/001_create_users_table.sql
      ```

3. **Generate Swagger documentation**:
   ```bash
   make swagger
   ```
   Or:
   ```bash
   swag init
   ```

4. **Run the application with hot reloading**:
   ```bash
   make run-local
   ```
   Or:
   ```bash
   air -c .air.toml
   ```

5. **Access the Swagger UI**:
   ```
   http://localhost:8080/api/v1/swagger/index.html
   ```

### Database Migrations

The database (`GO_CLEAN_ARCHITECTURE`) and schema (`users` table) are created automatically when the PostgreSQL container starts in Docker, thanks to the `POSTGRES_DB` environment variable and the migration script in `infrastructure/database/migrations/001_create_users_table.sql`. The migration is idempotent, so it can run multiple times without errors.

For local development or manual migration:
```bash
psql -U EngenMe -d GO_CLEAN_ARCHITECTURE -f infrastructure/database/migrations/001_create_users_table.sql
```

To add new migrations, create additional SQL files in `infrastructure/database/migrations` with numerical prefixes (e.g., `002_add_table.sql`).

### API Endpoints

The API is documented via Swagger at `http://localhost:8080/api/v1/swagger/index.html`. Key endpoints include:

#### Authentication
- `POST /api/v1/auth/login`: Authenticate a user and get a JWT token.
- `POST /api/v1/auth/signup`: Register a new user and get a JWT token.

#### Users
- `POST /api/v1/users`: Create a new user (public).
- `GET /api/v1/users`: Get all users (requires authentication).
- `GET /api/v1/users/:id`: Get user by ID (requires authentication).
- `GET /api/v1/users/email/:email`: Get user by email (public).
- `PUT /api/v1/users/:id`: Update user (requires authentication).
- `DELETE /api/v1/users/:id`: Delete user (requires authentication).

### Testing

Run unit tests (if any):
```bash
make test
```
Or:
```bash
go test ./... -v
```

**Note**: Add tests to the project in relevant directories (e.g., `services`, `handlers`) to ensure coverage.

### Makefile Commands

The `Makefile` simplifies common tasks:
```bash
make help           # Display available commands
make deps           # Install dependencies
make swagger        # Generate Swagger documentation
make build          # Build Docker images
make up             # Start application with Docker Compose
make down           # Stop and remove Docker containers
make clean          # Clean Docker resources and temporary files
make run-local      # Run application locally with Air
make test           # Run tests
make logs           # View Docker container logs
```

### Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/YourFeature`).
3. Commit changes (`git commit -m "Add YourFeature"`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

Please include tests and update documentation as needed.

### Troubleshooting

- **Database Issues**:
    - Verify the `GO_CLEAN_ARCHITECTURE` database and `users` table:
      ```bash
      docker exec -it go-clean-arch-postgres psql -U EngenMe -d GO_CLEAN_ARCHITECTURE -c "\dt"
      ```
    - Check PostgreSQL logs:
      ```bash
      docker logs go-clean-arch-postgres
      ```

- **Swagger Issues**:
    - The `Dockerfile` patches `docs/docs.go` to remove incompatible fields. If errors persist, ensure `swag@v1.16.3` for local runs:
      ```bash
      go install github.com/swaggo/swag/cmd/swag@v1.16.3
      ```

- **Connection Issues**:
    - Ensure `DB_HOST=postgres` in `.env` for Docker, or `DB_HOST=localhost` for local runs.
    - Check Docker network:
      ```bash
      docker network inspect go-clean-arch-network
      ```

### Notes

- **Database Persistence**: Data persists in the `postgres_data` Docker volume. Use `docker-compose down -v` to reset.
- **Security**: Do not commit `.env`. Generate a new `JWT_SECRET` for production (e.g., `uuidgen`).
- **Swagger UI**: Available in development (`ENV=development`). For production, add middleware to restrict access (see `routes.go`).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
# Go Clean Architecture Backend

A production-ready backend API built with Go, following Clean Architecture principles, CQRS pattern, and Mediator pattern.

## Technologies Used

- **Framework**: Gin Web Framework
- **ORM**: GORM with PostgreSQL
- **CQRS & Mediator**: Go-MediatR
- **Authentication**: JWT

## Features

- **Clean Architecture**: Separation of concerns with layers for handlers, services, repositories, and models
- **Generic CRUD**: Reusable CRUD operations that can be applied to any entity
- **User Management**: Full CRUD operations for User entity with email lookup
- **Authentication**: JWT-based authentication for securing API endpoints
- **Best Practices**: Error handling, logging, and environment-based configuration

## Project Structure

```
├── api                  # API Layer (Controllers, Middlewares, Routes)
├── application          # Application Layer (Commands, Queries, Services)
├── domain               # Domain Layer (Entities, Repository Interfaces)
├── infrastructure       # Infrastructure Layer (Database, External Services)
├── interfaces           # Interface Definitions
```

## Prerequisites

- Go 1.20 or higher
- PostgreSQL
- Docker (optional)

## Getting Started

### Environment Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-clean-architecture.git
cd go-clean-architecture
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Update the `.env` file with your PostgreSQL credentials and other settings.

### Running the Application

#### Without Docker

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run main.go
```

#### With Docker

1. Build and run using Docker:
```bash
docker build -t go-clean-architecture .
docker run -p 8080:8080 --env-file .env go-clean-architecture
```

### Database Migrations

The application automatically applies migrations when it starts. You can also find the SQL migration scripts in `infrastructure/database/migrations/`.

## API Endpoints

### Authentication
- `POST /api/v1/auth/login` - Login and get JWT token

### Users
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users` - Get all users (requires authentication)
- `GET /api/v1/users/:id` - Get user by ID (requires authentication)
- `GET /api/v1/users/email/:email` - Get user by email
- `PUT /api/v1/users/:id` - Update user (requires authentication)
- `DELETE /api/v1/users/:id` - Delete user (requires authentication)

## Testing

Run the tests with:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
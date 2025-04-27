FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Install build dependencies, Air, and swag CLI
RUN apk add --no-cache git && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Generate Swagger documentation 
RUN swag init

# Expose port
EXPOSE 8080

# Use Air for hot reloading
ENTRYPOINT ["air"]
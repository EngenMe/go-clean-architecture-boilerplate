FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Install build dependencies and Air for hot reloading
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Air will handle hot reloading
EXPOSE 8080
ENTRYPOINT ["air"]
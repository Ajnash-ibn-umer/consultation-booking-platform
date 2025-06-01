# Build the image:
#   docker build -t consultation-booking .
#
# Run the container:
#   docker run -d -p 8080:8080 --name consultation-app \
#   -e DB_HOST=your_db_host \
#   -e DB_USER=your_db_user \
#   -e DB_PASSWORD=your_db_password \
#   -e DB_NAME=your_db_name \
#   -e DB_PORT=5432 \
#   consultation-booking

# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Set working directory
WORKDIR /app


# Copy the .env file
COPY .env .

# Expose the port the app runs on
EXPOSE 8000

# Run the application
CMD ["./main"]
FROM golang:1.23-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Expose the application port
EXPOSE 8080

# Define build argument for mode (http or https)
ARG MODE=http

# Copy SSL certificate and key files if MODE is https
RUN if [ "$MODE" = "https" ]; then \
    echo "Running in HTTPS mode"; \
    cp cert.pem key.pem /app/; \
    fi

# Command to run the application
CMD if [ "$MODE" = "https" ]; then \
    /app/main -s -cert cert.pem -key key.pem; \
    else \
    /app/main; \
    fi
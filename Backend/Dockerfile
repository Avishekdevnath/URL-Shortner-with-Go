# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder  

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener ./cmd/server/main.go

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/url-shortener .

# Expose the port the app will run on
EXPOSE 8080

# Run the binary
CMD ["./url-shortener"]

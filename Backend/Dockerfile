# Step 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener ./cmd/server/main.go

# Step 2: Create the final runtime image
FROM alpine:latest

WORKDIR /app

# Install necessary packages for HTTPS connections
RUN apk --no-cache add ca-certificates

# Copy the Go binary from the builder stage
COPY --from=builder /app/url-shortener .

# Expose the server port
EXPOSE 8080

# Command to run the app
CMD ["./url-shortener"]
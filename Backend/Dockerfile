# S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\Dockerfile
# Step 1: Build the Go application
FROM golang:1.18 AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener .

# Step 2: Create the final runtime image
FROM alpine:latest

WORKDIR /root/

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Copy the Go binary from the builder stage
COPY --from=builder /app/url-shortener .

# Install Postgres client and Redis CLI
RUN apk add --no-cache postgresql-client redis

# Expose the server port
EXPOSE 8080

# Command to run the app
CMD ["./url-shortener"]

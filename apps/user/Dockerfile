# First stage: build the Go application
FROM golang:1.17-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for downloading dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fitsphere_user ./cmd/main.go

# Second stage: minimal runtime environment with only the compiled binary
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/fitsphere_user /app/

# Expose a port if your application serves on a specific port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/app/fitsphere_user"]

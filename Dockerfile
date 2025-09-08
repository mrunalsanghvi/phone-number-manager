# -------- STAGE 1: Build the Go binary --------
FROM golang:1.21 AS builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (for better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o server .

# -------- STAGE 2: Create a small final image --------
FROM alpine:latest

# Install certificates for HTTPS support (optional)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .

# Expose the port the app runs on
EXPOSE 8080

# Run the Go server
CMD ["./server"]

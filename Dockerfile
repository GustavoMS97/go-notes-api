FROM golang:1.23.0-alpine AS builder

# Install git (required by some Go packages)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o server cmd/main.go

# Stage 2: run
FROM alpine:latest

# Install CA certificates (required for HTTPS calls)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .

# Copy Swagger docs (optional if serving)
COPY --from=builder /app/docs ./docs

# Expose the port used by Fiber
EXPOSE 4000

# Command to run the executable
CMD ["./server"]

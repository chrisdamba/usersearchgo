# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go app
RUN go build -o usersearchgo main.go grpc_server.go

# Stage 2: Run the Go binary
FROM alpine:latest

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/usersearchgo /app/usersearchgo

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the binary
ENTRYPOINT ["/app/usersearchgo"]

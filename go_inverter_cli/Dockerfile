# Use a Go image to build the application
FROM golang:1.22 AS builder

# Set the target architecture for cross-compilation
ENV GOOS=linux
ENV GOARCH=386

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod .
COPY go.sum .

# Download all dependencies listed in go.mod
RUN go mod download

# Copy the Go application source code
COPY . .

# Ensure all modules are correctly resolved after copying source code
RUN go mod tidy

# Build the Go application
# CGO_ENABLED=0 is important for static linking, which makes the binary self-contained
# -o specifies the output file name
RUN CGO_ENABLED=0 go build -o inverter-cli .

# Use a minimal base image for the final runtime
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy mqtt.json into the image
COPY mqtt.json /app/mqtt.json

# Copy the compiled binary from the builder stage
COPY --from=builder /app/inverter-cli .

# Give execute permissions to the binary
RUN chmod +x /app/inverter-cli

# Set the entry point for the container
ENTRYPOINT ["/app/inverter-cli"]

FROM golang:1.24.4

# Set working directory
WORKDIR /app

# Install curl for downloading grpc-health-probe
RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

# Copy Go module files and download dependencies
COPY src/go.mod src/go.sum ./
RUN go mod tidy

# Copy the rest of the app
COPY src .

# Download and install grpc-health-probe
RUN curl -L -o /usr/local/bin/grpc-health-probe \
    https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.39/grpc_health_probe-linux-arm64 && \
    chmod +x /usr/local/bin/grpc-health-probe


# Build your Go app
RUN go build -o app cmd/main.go

# Expose gRPC port
EXPOSE 50051

# Start the app
CMD ["./app"]

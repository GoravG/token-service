# Token Service

A Go-based microservice for token management that integrates with Redis for token distribution and management.

## Overview

This service is designed to handle token generation and distribution through Redis pub/sub channels. It's built with Go and can be easily deployed using Docker.

## Prerequisites

- Go 1.24.2 or higher
- Docker (for containerized deployment)
- Redis server

## Configuration

The service can be configured using environment variables:

| Environment Variable | Description | Default Value |
|---------------------|-------------|---------------|
| `REDIS_ADDR` | Redis server address | `localhost:6379` |
| `REDIS_PASSWORD` | Redis server password | `""` (empty) |
| `PUBLISH_INTERVAL` | Token publishing interval in seconds | `5` |
| `REDIS_CHANNEL` | Redis channel name for publishing tokens | `tokens` |

## Project Structure

```
.
├── Dockerfile
├── go.mod
├── readme.md
└── internal/
    ├── cmd/
    │   └── server/
    │       └── main.go    # Main application entry point
    ├── config/
    │   └── config.go      # Configuration management
    └── utils/
        └── utils.go       # Utility functions
```

## Building and Running

### Local Development

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd token-service
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run internal/cmd/server/main.go
   ```

### Using Docker

1. Build the Docker image:
   ```bash
   docker build -t token-service .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 \
     -e REDIS_ADDR=redis:6379 \
     -e REDIS_PASSWORD=your_password \
     -e PUBLISH_INTERVAL=5 \
     -e REDIS_CHANNEL=tokens \
     token-service
   ```

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]

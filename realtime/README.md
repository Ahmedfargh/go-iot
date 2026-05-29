# Realtime Service

A specialized Go service for handling real-time notifications via WebSockets. It uses gRPC to communicate with the main `go-dashboard` service for user authentication.

## Features

- **WebSocket Support**: Real-time push notifications to connected clients.
- **gRPC Auth Client**: Communicates with the `go-dashboard` gRPC server to validate JWT tokens.
- **Database Integration**: Stores and retrieves notifications using GORM and MySQL.

## Tech Stack

- **Language**: Go
- **Framework**: [Gin](https://gin-gonic.com/)
- **gRPC**: [gRPC-Go](https://grpc.io/docs/languages/go/)
- **ORM**: [GORM](https://gorm.io/)
- **Communication**: WebSockets

## Getting Started

### Environment Setup

1. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```
2. Update the environment variables in `.env`. Ensure the database credentials match your setup.

### Prerequisites

The `go-dashboard` service must be running and its gRPC server should be listening (default port `50051`) for authentication to work.

### Installation

```bash
go mod download
```

### Running the Application

```bash
go run realtime.go
```

The service will start on port `8081` by default.

## API & WebSockets

- **WebSocket Endpoint**: `ws://localhost:8081/ws`
- **Authentication**: Requires a `Bearer` token in the `Authorization` header during the initial HTTP handshake or as a query parameter (depending on client implementation). The middleware validates this token via gRPC.

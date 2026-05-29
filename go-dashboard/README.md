# Go Dashboard API

The backend service for the IoT management system, providing RESTful APIs for device management, user authentication, and system monitoring.

## Features

- **Auth**: JWT-based authentication for admins and users.
- **Device Management**: CRUD operations for IoT devices.
- **RBAC**: Role-based access control with permissions.
- **Real-time**: WebSocket integration and gRPC server for cross-service authentication.
- **Messaging**: RabbitMQ integration for asynchronous tasks.
- **Database**: GORM with MySQL support and MongoDB integration.
- **gRPC Server**: Provides authentication services to other microservices (e.g., `realtime`).

## Tech Stack

- **Language**: Go
- **Framework**: [Gin](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: MySQL, MongoDB
- **Queue**: RabbitMQ
- **Middleware**: JWT Auth, Rate Limiting, Permission Checks

## Getting Started

### Environment Setup

1. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```
2. Update the environment variables in `.env` with your database and RabbitMQ credentials.

### Installation

```bash
go mod download
```

### Running the Application

```bash
go run go-dashboard.go
```

## API Documentation

A Postman collection is available at `postman_collection.json` for testing the APIs.

# Go IoT System

A comprehensive IoT management system consisting of a Go-based backend and a Vue-based frontend dashboard.

## Project Architecture

The system is composed of three main components:

- **`go-dashboard/`**: The core backend API service.
  - RESTful APIs for device management, user authentication, and RBAC.
  - gRPC server for cross-service authentication.
  - Integration with MySQL, MongoDB, and RabbitMQ.
- **`realtime/`**: A specialized service for real-time notifications via WebSockets.
  - Authenticates requests by communicating with `go-dashboard` via gRPC.
  - Manages live notification streams for users.
- **`vue-dashboard/`**: A modern, responsive frontend dashboard.
  - Built with Vue 3, Vite, Pinia, and Tailwind CSS.
  - Provides a visual interface for monitoring and managing the IoT ecosystem.

## Prerequisites

- [Go](https://go.dev/doc/install) (v1.23+)
- [Node.js](https://nodejs.org/) (v18+)
- [MySQL](https://www.mysql.com/)
- [MongoDB](https://www.mongodb.com/)
- [RabbitMQ](https://www.rabbitmq.com/)

## Getting Started

### 1. Clone the repository

```bash
git clone <repository-url>
cd go-iot
```

### 2. Set up the Core Backend (go-dashboard)

```bash
cd go-dashboard
cp .env.example .env
# Update .env with your credentials
go mod download
go run go-dashboard.go
```

### 3. Set up the Realtime Service

```bash
cd ../realtime
cp .env.example .env
# Update .env with your credentials
go mod download
go run realtime.go
```

### 4. Set up the Frontend Dashboard

```bash
cd ../vue-dashboard
npm install
npm run dev
```

## Infrastructure

To simplify setup, you can use the following default ports:

- **go-dashboard**: `http://localhost:8080`
- **gRPC Auth**: `localhost:50051`
- **gRPC Notification**: `localhost:50052`
- **realtime**: `http://localhost:8081` (WebSockets)
- **vue-dashboard**: `http://localhost:5173`

## License

[MIT](LICENSE)

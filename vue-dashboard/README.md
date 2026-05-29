# Vue Dashboard Frontend

The frontend dashboard for the Go IoT management system, providing a modern interface for managing and monitoring IoT devices.

## Features

- **System Overview**: High-level metrics and system health monitoring.
- **Device Management**: Interface for adding, updating, and monitoring IoT devices.
- **Admin Management**: Comprehensive management of users, roles, and permissions.
- **Real-time Updates**: Integrates with the `realtime` service for live notifications.
- **Responsive Design**: Built with Tailwind CSS 4 for a seamless experience across devices.
- **State Management**: Centralized state management using Pinia.

## Tech Stack

- **Framework**: [Vue 3](https://vuejs.org/)
- **Build Tool**: [Vite](https://vitejs.dev/)
- **Styling**: [Tailwind CSS 4](https://tailwindcss.com/)
- **State Management**: [Pinia](https://pinia.vuejs.org/)
- **Routing**: [Vue Router](https://router.vuejs.org/)
- **HTTP Client**: [Axios](https://axios-http.com/)

## Getting Started

### Installation

```bash
npm install
```

### Development

```bash
npm run dev
```

### Build

```bash
npm run build
```

## Configuration

The frontend connects to the backend API. Ensure the API URL is correctly configured in `src/api/client.js` or through environment variables if applicable.

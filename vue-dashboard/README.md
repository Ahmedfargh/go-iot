# Vue Dashboard Frontend

A modern, responsive dashboard for managing and monitoring IoT devices.

## Features

- **System Overview**: High-level metrics and system health.
- **Device Management**: Interface for adding and monitoring IoT devices.
- **Admin Management**: Manage users, roles, and permissions.
- **Responsive Design**: Built with Tailwind CSS for various screen sizes.
- **State Management**: Centralized state using Pinia.

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

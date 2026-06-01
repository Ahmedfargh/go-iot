#!/bin/bash

set -e

echo "========================================="
echo "Go-IoT Project Installation Script"
echo "========================================="
echo ""

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$PROJECT_ROOT"

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[✓]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

# Check prerequisites
log_info "Checking prerequisites..."

if ! command -v go &> /dev/null; then
    log_error "Go is not installed. Please install Go 1.18+"
    exit 1
fi
log_success "Go $(go version | awk '{print $3}')"

if ! command -v node &> /dev/null; then
    log_error "Node.js is not installed. Please install Node.js 16+"
    exit 1
fi
log_success "Node.js $(node --version)"

if ! command -v npm &> /dev/null; then
    log_error "npm is not installed. Please install npm"
    exit 1
fi
log_success "npm $(npm --version)"

echo ""

# Install go-dashboard
log_info "Installing go-dashboard service..."
cd "$PROJECT_ROOT/go-dashboard"

if [ ! -f "go.mod" ]; then
    log_error "go.mod not found in go-dashboard"
    exit 1
fi

log_info "Downloading Go dependencies..."
go mod download
go mod tidy
log_success "Go dependencies downloaded"

log_info "Building go-dashboard..."
go build -o go-dashboard .
log_success "go-dashboard built successfully"

echo ""

# Install realtime
log_info "Installing realtime service..."
cd "$PROJECT_ROOT/realtime"

if [ ! -f "go.mod" ]; then
    log_error "go.mod not found in realtime"
    exit 1
fi

log_info "Downloading Go dependencies..."
go mod download
go mod tidy
log_success "Go dependencies downloaded"

log_info "Building realtime..."
go build -o realtime .
log_success "realtime built successfully"

echo ""

# Install vue-dashboard
log_info "Installing vue-dashboard..."
cd "$PROJECT_ROOT/vue-dashboard"

if [ ! -f "package.json" ]; then
    log_error "package.json not found in vue-dashboard"
    exit 1
fi

log_info "Installing Node.js dependencies..."
npm install
log_success "Node.js dependencies installed"

log_info "Building Vue.js dashboard..."
npm run build
log_success "vue-dashboard built successfully"

echo ""
echo "========================================="
log_success "All services installed successfully!"
echo "========================================="
echo ""
log_info "Next steps:"
echo ""
echo "1. Create .env files in each service directory:"
echo "   - go-dashboard/.env"
echo "   - realtime/.env"
echo ""
echo "2. Configure environment variables (see README files)"
echo ""
echo "3. Start the services:"
echo "   # Terminal 1: go-dashboard"
echo "   cd go-dashboard && ./go-dashboard"
echo ""
echo "   # Terminal 2: realtime"
echo "   cd realtime && ./realtime"
echo ""
echo "   # Terminal 3: vue-dashboard (dev server)"
echo "   cd vue-dashboard && npm run dev"
echo ""
echo "4. Access:"
echo "   - Dashboard: http://localhost:5173"
echo "   - API: http://localhost:8080"
echo "   - WebSocket: ws://localhost:8081"
echo "   - gRPC Auth: localhost:50051"
echo ""

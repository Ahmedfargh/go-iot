#!/bin/bash

set -e

echo "========================================="
echo "realtime Installation Script"
echo "========================================="
echo ""

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$PROJECT_DIR"

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[✓]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check prerequisites
log_info "Checking prerequisites..."

if ! command -v go &> /dev/null; then
    log_error "Go is not installed. Please install Go 1.18+"
    exit 1
fi
log_success "Go $(go version | awk '{print $3}')"

if [ ! -f "go.mod" ]; then
    log_error "go.mod not found in $(pwd)"
    exit 1
fi

echo ""
log_info "Downloading Go dependencies..."
go mod download
go mod tidy
log_success "Go dependencies downloaded"

echo ""
log_info "Building realtime..."
go build -o realtime .
log_success "realtime built successfully"

echo ""
echo "========================================="
log_success "Installation complete!"
echo "========================================="
echo ""
log_info "To start the service:"
echo "  ./realtime"
echo ""
log_info "Or with custom port:"
echo "  export REALTIME_HTTP_PORT=8081 && ./realtime"
echo ""

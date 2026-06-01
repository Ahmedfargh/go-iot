#!/bin/bash

set -e

echo "========================================="
echo "vue-dashboard Installation Script"
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

if [ ! -f "package.json" ]; then
    log_error "package.json not found in $(pwd)"
    exit 1
fi

echo ""
log_info "Installing Node.js dependencies..."
npm install
log_success "Node.js dependencies installed"

echo ""
log_info "Building Vue.js dashboard..."
npm run build
log_success "vue-dashboard built successfully"

echo ""
echo "========================================="
log_success "Installation complete!"
echo "========================================="
echo ""
log_info "To start development server:"
echo "  npm run dev"
echo ""
log_info "Production build is in: ./dist/"
echo ""

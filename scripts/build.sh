#!/bin/bash

# Set the application name
APP_NAME="vin-search"

# Create output directories
mkdir -p bin

# Function to build the application
build() {
    local os=$1
    local arch=$2
    local output="${APP_NAME}_${os}_${arch}"

    echo "Building for ${os} (${arch})..."
    GOOS=${os} GOARCH=${arch} go build -o bin/${output} cmd/vin-search/main.go

    if [ $? -ne 0 ]; then
        echo "Build failed for ${os} (${arch})"
        exit 1
    else
        echo "Build succeeded for ${os} (${arch})"
    fi
}

# Build for macOS (Intel)
build darwin amd64

# Build for macOS (ARM)
build darwin arm64

# Build for Ubuntu (Linux)
build linux amd64

echo "Build complete. Binaries are located in the bin/ directory."

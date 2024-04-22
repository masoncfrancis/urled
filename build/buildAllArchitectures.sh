#!/bin/bash

# Check if project directory and version number are provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <project_directory> <version>"
    exit 1
fi

PROJECT_DIR="$1"
VERSION="$2"
PROJECT_NAME=$(basename "$PROJECT_DIR")

# Ensure the project directory exists
if [ ! -d "$PROJECT_DIR" ]; then
    echo "Project directory not found: $PROJECT_DIR"
    exit 1
fi

# Build for darwin/386
env GOOS=darwin GOARCH=386 go build -o "./${PROJECT_NAME}_darwin_386_v${VERSION}" "$PROJECT_DIR"

# Build for darwin/amd64
env GOOS=darwin GOARCH=amd64 go build -o "./${PROJECT_NAME}_darwin_amd64_v${VERSION}" "$PROJECT_DIR"

# Build for darwin/arm
env GOOS=darwin GOARCH=arm go build -o "./${PROJECT_NAME}_darwin_arm_v${VERSION}" "$PROJECT_DIR"

# Build for darwin/arm64
env GOOS=darwin GOARCH=arm64 go build -o "./${PROJECT_NAME}_darwin_arm64_v${VERSION}" "$PROJECT_DIR"

# Build for freebsd/386
env GOOS=freebsd GOARCH=386 go build -o "./${PROJECT_NAME}_freebsd_386_v${VERSION}" "$PROJECT_DIR"

# Build for freebsd/amd64
env GOOS=freebsd GOARCH=amd64 go build -o "./${PROJECT_NAME}_freebsd_amd64_v${VERSION}" "$PROJECT_DIR"

# Build for freebsd/arm
env GOOS=freebsd GOARCH=arm go build -o "./${PROJECT_NAME}_freebsd_arm_v${VERSION}" "$PROJECT_DIR"

# Build for linux/386
env GOOS=linux GOARCH=386 go build -o "./${PROJECT_NAME}_linux_386_v${VERSION}" "$PROJECT_DIR"

# Build for linux/amd64
env GOOS=linux GOARCH=amd64 go build -o "./${PROJECT_NAME}_linux_amd64_v${VERSION}" "$PROJECT_DIR"

# Build for linux/armv5
env GOOS=linux GOARCH=arm GOARM=5 go build -o "./${PROJECT_NAME}_linux_armv5_v${VERSION}" "$PROJECT_DIR"

# Build for linux/armv6
env GOOS=linux GOARCH=arm GOARM=6 go build -o "./${PROJECT_NAME}_linux_armv6_v${VERSION}" "$PROJECT_DIR"

# Build for linux/armv7
env GOOS=linux GOARCH=arm GOARM=7 go build -o "./${PROJECT_NAME}_linux_armv7_v${VERSION}" "$PROJECT_DIR"

# Build for linux/arm64
env GOOS=linux GOARCH=arm64 go build -o "./${PROJECT_NAME}_linux_arm64_v${VERSION}" "$PROJECT_DIR"

# Build for windows/386
env GOOS=windows GOARCH=386 go build -o "./${PROJECT_NAME}_windows_386_v${VERSION}.exe" "$PROJECT_DIR"

# Build for windows/amd64
env GOOS=windows GOARCH=amd64 go build -o "./${PROJECT_NAME}_windows_amd64_v${VERSION}.exe" "$PROJECT_DIR"

# Build for windows/arm
env GOOS=windows GOARCH=arm go build -o "./${PROJECT_NAME}_windows_arm_v${VERSION}.exe" "$PROJECT_DIR"

echo "Builds completed."

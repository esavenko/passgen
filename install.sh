#!/bin/bash

set -e

# Check OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
    linux)
        OS="Linux"
        ;;
    darwin)
        OS="Darwin"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

case $ARCH in
    x86_64|amd64)
        ARCH="x86_64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Get actual version
LATEST_VERSION=$(curl -s https://api.github.com/repos/esavenko/passgen/releases/latest | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo "Failed to get latest version"
    exit 1
fi

# URL
DOWNLOAD_URL="https://github.com/esavenko/passgen/releases/download/${LATEST_VERSION}/passgen_${OS}_${ARCH}.tar.gz"

# Dir for installation
INSTALL_DIR="/usr/local/bin"
if [ "$(id -u)" -ne 0 ] && [ ! -w "$INSTALL_DIR" ]; then
    INSTALL_DIR="$HOME/bin"
    mkdir -p "$INSTALL_DIR"
    echo "Note: Add $INSTALL_DIR to your PATH if it's not already there."
fi

# Download and install
echo "Downloading passgen ${LATEST_VERSION} for ${OS} ${ARCH}..."
curl -L -o /tmp/passgen.tar.gz "$DOWNLOAD_URL"

echo "Extracting..."
tar -xzf /tmp/passgen.tar.gz -C /tmp

echo "Installing to $INSTALL_DIR..."
mv /tmp/passgen "$INSTALL_DIR/passgen"
chmod +x "$INSTALL_DIR/passgen"

rm /tmp/passgen.tar.gz

echo "Installation complete! Run 'passgen' to start."
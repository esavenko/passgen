#!/bin/bash

set -e

# Определяем ОС и архитектуру
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

# Получаем последнюю версию из GitHub API
LATEST_VERSION=$(curl -s https://api.github.com/repos/esavenko/passgen/releases/latest | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo "Failed to get latest version"
    exit 1
fi

# URL для скачивания
DOWNLOAD_URL="https://github.com/esavenko/passgen/releases/download/${LATEST_VERSION}/passgen_${OS}_${ARCH}.tar.gz"

# Директория для установки
INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
    INSTALL_DIR="$HOME/bin"
    mkdir -p "$INSTALL_DIR"
    export PATH="$INSTALL_DIR:$PATH"
fi

# Скачиваем и устанавливаем
echo "Downloading passgen ${LATEST_VERSION} for ${OS} ${ARCH}..."
curl -L -o /tmp/passgen.tar.gz "$DOWNLOAD_URL"

echo "Extracting..."
tar -xzf /tmp/passgen.tar.gz -C /tmp

echo "Installing to $INSTALL_DIR..."
mv /tmp/passgen "$INSTALL_DIR/passgen"
chmod +x "$INSTALL_DIR/passgen"

rm /tmp/passgen.tar.gz

echo "Installation complete! Run 'passgen' to start."
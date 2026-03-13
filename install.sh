#!/bin/bash

set -euo pipefail

REPO="esavenko/passgen"
BINARY="passgen"

# --- Detect OS ---
OS="$(uname -s)"
case "$OS" in
    Linux)  OS="Linux" ;;
    Darwin) OS="Darwin" ;;
    *)
        echo "Error: unsupported OS: $OS" >&2
        exit 1
        ;;
esac

# --- Detect architecture ---
ARCH="$(uname -m)"
case "$ARCH" in
    x86_64|amd64)   ARCH="x86_64" ;;
    arm64|aarch64)   ARCH="arm64" ;;
    *)
        echo "Error: unsupported architecture: $ARCH" >&2
        exit 1
        ;;
esac

# --- Fetch latest version ---
LATEST_VERSION="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')"

if [ -z "$LATEST_VERSION" ]; then
    echo "Error: failed to determine latest version" >&2
    exit 1
fi

echo "Installing ${BINARY} ${LATEST_VERSION} (${OS}/${ARCH})..."

# --- Prepare temp directory ---
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

ARCHIVE="${BINARY}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_VERSION}/${ARCHIVE}"
CHECKSUMS_URL="https://github.com/${REPO}/releases/download/${LATEST_VERSION}/${BINARY}_${LATEST_VERSION#v}_checksums.txt"

# --- Download archive and checksums ---
curl -fsSL -o "${TMP_DIR}/${ARCHIVE}" "$DOWNLOAD_URL"
curl -fsSL -o "${TMP_DIR}/checksums.txt" "$CHECKSUMS_URL"

# --- Verify checksum ---
EXPECTED="$(grep "${ARCHIVE}" "${TMP_DIR}/checksums.txt" | awk '{print $1}')"
if [ -z "$EXPECTED" ]; then
    echo "Warning: checksum not found for ${ARCHIVE}, skipping verification" >&2
else
    if command -v sha256sum &>/dev/null; then
        ACTUAL="$(sha256sum "${TMP_DIR}/${ARCHIVE}" | awk '{print $1}')"
    else
        ACTUAL="$(shasum -a 256 "${TMP_DIR}/${ARCHIVE}" | awk '{print $1}')"
    fi

    if [ "$EXPECTED" != "$ACTUAL" ]; then
        echo "Error: checksum mismatch" >&2
        echo "  expected: ${EXPECTED}" >&2
        echo "  actual:   ${ACTUAL}" >&2
        exit 1
    fi
    echo "Checksum verified."
fi

# --- Extract ---
tar -xzf "${TMP_DIR}/${ARCHIVE}" -C "$TMP_DIR"

# --- Install ---
INSTALL_DIR="/usr/local/bin"
if [ "$(id -u)" -ne 0 ] && [ ! -w "$INSTALL_DIR" ]; then
    INSTALL_DIR="${HOME}/.local/bin"
    mkdir -p "$INSTALL_DIR"
elif [ ! -d "$INSTALL_DIR" ]; then
    mkdir -p "$INSTALL_DIR"
fi

mv "${TMP_DIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
chmod +x "${INSTALL_DIR}/${BINARY}"

echo "Installed ${BINARY} to ${INSTALL_DIR}/${BINARY}"

# --- Check PATH ---
case ":${PATH}:" in
    *":${INSTALL_DIR}:"*) ;;
    *)
        echo ""
        echo "Note: ${INSTALL_DIR} is not in your PATH."
        echo "Add it by running:"
        echo "  export PATH=\"${INSTALL_DIR}:\$PATH\""
        ;;
esac

echo "Done! Run '${BINARY}' to get started."

#!/usr/bin/bash
set -euo pipefail

REPO="Fadilix/couik"
INSTALL_DIR="/usr/local/bin"

die() { echo "error: $*" >&2; exit 1; }

case "$(uname -m)" in
    x86_64|amd64)  ARCH="x86_64" ;;
    aarch64|arm64) ARCH="arm64"  ;;
    *) die "unsupported architecture: $(uname -m)" ;;
esac

if [[ -f /etc/arch-release ]] || grep -qi "arch" /etc/os-release 2>/dev/null; then
    for helper in yay paru trizen pikaur; do
        if command -v "$helper" &>/dev/null; then
            echo "Arch Linux detected — installing via $helper..."
            "$helper" -S --noconfirm couik-bin
            exit 0
        fi
    done
fi

ASSET="couik_Linux_${ARCH}.tar.gz"
VERSION=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')
URL="https://github.com/${REPO}/releases/download/${VERSION}/${ASSET}"

echo "Installing couik ${VERSION} (${ARCH})..."

TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

curl -fsSL --progress-bar -o "${TMP}/${ASSET}" "$URL"
tar -xzf "${TMP}/${ASSET}" -C "$TMP"

[[ -f "${TMP}/couik" ]] || die "binary not found in archive"

if [[ -w "$INSTALL_DIR" ]]; then
    install -m755 "${TMP}/couik" "${INSTALL_DIR}/couik"
else
    sudo install -m755 "${TMP}/couik" "${INSTALL_DIR}/couik"
fi

echo "Done. Run: couik"

#!/usr/bin/env sh
set -eu

REPO="KlarLang/loom"

VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest \
    | grep '"tag_name"' \
    | cut -d '"' -f 4)

if [ -z "$VERSION" ]; then
    echo "❌ Could not fetch latest version."
    exit 1
fi

echo "Latest version: $VERSION"

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux)  OS="Linux" ;;
  Darwin) OS="Darwin" ;;
  MINGW*|MSYS*|CYGWIN*) OS="Windows" ;;
  *) echo "❌ Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64) ARCH="x86_64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  i386|i686) ARCH="i386" ;;
  *) echo "❌ Unsupported arch: $ARCH"; exit 1 ;;
esac

FILE_NAME="loom_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILE_NAME"

echo "Downloading: $FILE_NAME"

TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

curl -L --fail "$URL" -o "$TMPDIR/$FILE_NAME"

tar -xzf "$TMPDIR/$FILE_NAME" -C "$TMPDIR"

LOOM_BIN="$(find "$TMPDIR" -type f -name loom -print | head -n 1 || true)"

if [ -z "$LOOM_BIN" ]; then
  echo "❌ loom binary not found inside archive."
  exit 1
fi

chmod +x "$LOOM_BIN"

INSTALL_PATH="$(command -v loom || true)"

if [ -z "$INSTALL_PATH" ]; then
    echo "loom not installed — installing fresh."

    if command -v sudo >/dev/null 2>&1; then
        sudo mv "$LOOM_BIN" /usr/local/bin/loom
    else
        mkdir -p "$HOME/.local/bin"
        mv "$LOOM_BIN" "$HOME/.local/bin/loom"
    fi

    echo "✔ Installed loom $VERSION"
    exit 0
fi

echo "Installing over: $INSTALL_PATH"

if [ -w "$INSTALL_PATH" ]; then
    mv "$LOOM_BIN" "$INSTALL_PATH"
else
    sudo mv "$LOOM_BIN" "$INSTALL_PATH"
fi

echo "✔ Updated loom to $VERSION"
echo "→ Run: loom --version"

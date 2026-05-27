set -e

echo "Buildando Loom"
zig build

echo "Copiando pro path global"
cp zig-out/bin/loom /usr/local/bin/
echo "Copiado ✓"

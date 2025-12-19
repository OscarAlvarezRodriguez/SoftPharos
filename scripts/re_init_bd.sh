#!/bin/bash
set -e  # detiene el script si ocurre algún error

echo "🚀 Reiniciando entorno de PostgreSQL con Docker Compose..."

# Nombre del contenedor y volumen según tu docker-compose.yml
CONTAINER_NAME="pg-demo-compose"
VOLUME_NAME="softpharos_pgdata"

echo "🛑 Deteniendo y eliminando contenedor: $CONTAINER_NAME (si existe)..."
docker stop $CONTAINER_NAME 2>/dev/null || true
docker rm $CONTAINER_NAME 2>/dev/null || true

echo "🧹 Eliminando volumen de datos: $VOLUME_NAME (si existe)..."
docker volume rm $VOLUME_NAME 2>/dev/null || true

echo "🧽 Limpiando volúmenes huérfanos..."
docker volume prune -f

echo "🧰 Reconstruyendo entorno limpio..."
# Volver al directorio raíz y ejecutar setup.sh
cd "$(dirname "$0")/.."
bash scripts/setup.sh

echo "✅ Base de datos PostgreSQL reiniciada y contenedor en ejecución."

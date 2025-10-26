#!/bin/bash
set -e

echo "🔧 Iniciando setup del proyecto..."

# --- 0. Cargar variables del entorno (.env) ---
if [ -f .env ]; then
  echo "🕐 Cargando variables desde .env..."
  set -a
  source .env
  set +a
else
  echo "❌️ No se encontró el archivo .env. Asegúrate de tener uno antes de continuar."
  exit 1
fi

# --- 1. Levantar base de datos con Docker Compose ---
echo "🐘 Levantando PostgreSQL con Docker Compose..."
if command -v docker-compose >/dev/null 2>&1; then
  docker-compose up -d
else
  docker compose up -d
fi

PG_CONTAINER=pg-demo-compose

if [ -z "$PG_CONTAINER" ]; then
  echo "❌ No se encontró un contenedor PostgreSQL corriendo. Verifica tu docker-compose.yml"
  exit 1
fi

# --- 2. Esperar a que PostgreSQL esté listo ---
echo "🕐 Esperando a que PostgreSQL acepte conexiones..."
until docker exec "$PG_CONTAINER" pg_isready -U "$DB_USER" -d "$DB_NAME" >/dev/null 2>&1; do
  sleep 1
done
echo "✅ PostgreSQL está listo."

# --- 3. Ejecutar script SQL de inicialización (solo si la BD está vacía) ---

INIT_SQL="Proyecto/Backend/cmd/bd/init.sql"
SEED_SQL="Proyecto/Backend/cmd/bd/seed.sql"
SEED_DEV_SQL="Proyecto/Backend/cmd/bd/seed_dev.sql"

if [ -f "$INIT_SQL" ]; then
  echo "🔍 Verificando si la base de datos ya fue inicializada..."
  TABLE_COUNT=$(docker exec -i "$PG_CONTAINER" psql -U "$DB_USER" -d "$DB_NAME" -t -c \
    "SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public';" | tr -d '[:space:]')

  if [ "$TABLE_COUNT" = "0" ] || [ -z "$TABLE_COUNT" ]; then
    echo "🕐 Ejecutando script SQL de inicialización..."
    docker exec -i "$PG_CONTAINER" psql -U "$DB_USER" -d "$DB_NAME" < "$INIT_SQL"

    # Ejecutar seed de datos esenciales
    if [ -f "$SEED_SQL" ]; then
      echo "🌱 Poblando base de datos con datos esenciales (seed.sql)..."
      docker exec -i "$PG_CONTAINER" psql -U "$DB_USER" -d "$DB_NAME" < "$SEED_SQL"
    fi

    # Ejecutar seed de desarrollo (solo si existe)
    if [ -f "$SEED_DEV_SQL" ]; then
      echo "🔍 ¿Deseas cargar datos de desarrollo/testing? (y/N)"
      read -r LOAD_DEV_DATA
      if [[ "$LOAD_DEV_DATA" =~ ^[Yy]$ ]]; then
        echo "🌱 Poblando base de datos con datos de desarrollo (seed_dev.sql)..."
        docker exec -i "$PG_CONTAINER" psql -U "$DB_USER" -d "$DB_NAME" < "$SEED_DEV_SQL"
      else
        echo "⏭️ Saltando datos de desarrollo."
      fi
    fi
  else
    echo "✅ Base de datos ya inicializada. No se ejecutará init.sql."
    echo "💡 Para poblar datos manualmente, ejecuta:"
    echo "   docker exec -i $PG_CONTAINER psql -U $DB_USER -d $DB_NAME < $SEED_SQL"
  fi
else
  echo "⚠️ No se encontró el archivo $INIT_SQL"
fi

# --- 4. Backend (Go) ---
echo "🔍 Verificando dependencias del Backend..."
cd Proyecto/Backend || exit

if ! command -v go >/dev/null 2>&1; then
  echo "❌ Go no está instalado. Instálalo antes de continuar."
  exit 1
fi

if [ ! -f "go.mod" ]; then
  echo "❌️ No se encontró go.mod. Ejecuta 'go mod init <nombre>' manualmente."
else
  echo "🕐 Ejecutando 'go mod tidy'..."
  go mod tidy
fi

cd ../../

# --- 5. Frontend (Vue 3) ---
echo "🔍 Verificando dependencias del Frontend (Vue 3)..."
cd Proyecto/Frontend || exit

if ! command -v npm >/dev/null 2>&1; then
  echo "❌ npm no está instalado. Instálalo antes de continuar."
  exit 1
fi

if [ -d "node_modules" ]; then
  echo "✅ Dependencias ya instaladas. Saltando npm install."
else
  echo "🔍 Instalando dependencias..."
  npm install
fi
echo "✅ Frontend verificado correctamente."

cd ../../
echo "✅ Setup finalizado correctamente."

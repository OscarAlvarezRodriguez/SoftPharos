#!/bin/bash

echo "🔧 Iniciando setup del proyecto..."

# 1. Levantar base de datos con Docker Compose
echo "🐘 Levantando PostgreSQL con Docker Compose..."
if command -v docker-compose >/dev/null 2>&1; then
  docker-compose up -d
else
  docker compose up -d
fi

# Esperar unos segundos a que PostgreSQL esté listo
echo "⏳ Esperando a que la base de datos esté lista..."
sleep 5

# 2. Ejecutar el script SQL de inicialización
if [ -f "Proyecto/Backend/cmd/bd/init.sql" ]; then
  echo "📄 Ejecutando script SQL de inicialización..."
  docker exec -i pg-demo-compose psql -U testuser -d testdb < Proyecto/Backend/cmd/bd/init.sql
else
  echo "⚠️ No se encontró init.sql en Proyecto/Backend/cmd/bd/"
fi

# 3. Backend (Go)
echo "📦 Configurando Backend en Go..."
cd Proyecto/Backend || exit
if [ ! -f "go.mod" ]; then
  go mod init backend
fi
go mod tidy
cd ../../

# 4. Frontend (opcional)
#read -p "¿Quieres usar Vue o React? (vue/react): " choice
#
#if [ "$choice" = "vue" ]; then
#  echo "📦 Creando frontend con Vue..."
#  cd Proyecto || exit
#  npm create vue@latest Frontend
#  cd ..
#elif [ "$choice" = "react" ]; then
#  echo "📦 Creando frontend con React..."
#  cd Proyecto || exit
#  npx create-react-app Frontend
#  cd ..
#else
#  echo "⚠️ Opción no válida. No se instaló frontend."
#fi

echo "✅ Setup finalizado."

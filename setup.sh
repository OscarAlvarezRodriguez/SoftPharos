#!/bin/bash

echo "🔧 Iniciando setup del proyecto..."

# Backend (Go)
echo "📦 Instalando dependencias de Go..."
cd Proyecto/Backend || exit
go mod init backend
go get
go mod tidy
cd ../../

# Frontend
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

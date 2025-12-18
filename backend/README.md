# Backend - SoftPharos API

API REST desarrollada en Go con arquitectura hexagonal (Clean Architecture).

## 🏗️ Arquitectura

```
backend/
├── cmd/
│   ├── app/              # Configuración de rutas
│   ├── buildingAPI/      # Inyección de dependencias
│   └── bd/               # Scripts SQL
├── internal/
│   ├── controllers/      # Handlers HTTP
│   ├── core/
│   │   ├── domain/       # Entidades de negocio
│   │   ├── ports/        # Interfaces (contratos)
│   │   ├── repository/   # Implementación repositorios
│   │   └── services/     # Lógica de negocio
│   └── infra/
│       └── databases/    # PostgreSQL + GORM
└── main.go
```

## 🚀 Ejecución

```bash
# Desarrollo
go run main.go

# Producción
go build -o softpharos
./softpharos
```

## 🧪 Tests

```bash
# Todos los tests
go test ./...

# Con cobertura
bash run_tests.sh
```

## 📦 Dependencias principales

- **Gin**: Framework web
- **GORM**: ORM para PostgreSQL
- **godotenv**: Variables de entorno

## 🔌 API

- Puerto: `8080` (configurable en `.env`)
- CORS habilitado para desarrollo
- Formato respuestas: JSON

## 🌍 Variables de entorno

Crear archivo `.env` en la raíz del proyecto:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=softpharos
PORT=8080
ENV=development
```

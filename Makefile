.PHONY: help setup dev-backend dev-frontend test lint build clean db-reset

help: ## Muestra esta ayuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

setup: ## Configuración inicial completa del proyecto
	@bash scripts/setup.sh

dev-backend: ## Inicia el backend en modo desarrollo (puerto 8080)
	@cd backend && go run main.go

dev-frontend: ## Inicia el frontend en modo desarrollo (puerto 5173)
	@cd frontend && npm run dev

test: ## Ejecuta todos los tests (backend + frontend)
	@echo "🧪 Ejecutando tests del backend..."
	@cd backend && go test ./...
	@echo "🧪 Ejecutando tests del frontend..."
	@cd frontend && npm run test:unit

lint: ## Ejecuta linters en todo el proyecto
	@echo "🔍 Ejecutando linters del backend..."
	@cd backend && go fmt ./...
	@echo "🔍 Ejecutando linters del frontend..."
	@cd frontend && npm run lint

build: ## Compila todo el proyecto (backend + frontend)
	@echo "🔨 Compilando backend..."
	@cd backend && go build -o ../bin/softpharos main.go
	@echo "🔨 Compilando frontend..."
	@cd frontend && npm run build
	@echo "✅ Build completado: bin/softpharos (backend) y frontend/dist/ (frontend)"

db-reset: ## Reinicia la base de datos desde cero
	@bash scripts/re_init_bd.sh

clean: ## Limpia archivos generados y temporales
	@echo "🧹 Limpiando archivos temporales..."
	@cd backend && go clean
	@rm -rf bin/
	@rm -rf frontend/dist/
	@rm -rf frontend/node_modules/.vite
	@echo "✅ Limpieza completada"

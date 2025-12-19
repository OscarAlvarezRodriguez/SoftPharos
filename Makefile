.PHONY: help setup dev dev-backend dev-frontend test lint build clean db-reset

help: ## Muestra esta ayuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

setup: ## Configuración inicial completa del proyecto
	@bash scripts/setup.sh

dev: ## Muestra cómo iniciar el entorno de desarrollo
	@echo "🚀 Para desarrollo, abre DOS terminales y ejecuta:"
	@echo ""
	@echo "  Terminal 1:  make dev-backend"
	@echo "  Terminal 2:  make dev-frontend"
	@echo ""
	@echo "Luego accede a:"
	@echo "  • Frontend: http://localhost:5173"
	@echo "  • Backend:  http://localhost:8080"

dev-backend: ## Inicia el backend en modo desarrollo (puerto 8080)
	@echo "🚀 Iniciando backend en http://localhost:8080"
	@cd backend && go run main.go

dev-frontend: ## Inicia el frontend en modo desarrollo (puerto 5173)
	@echo "🚀 Iniciando frontend en http://localhost:5173"
	@cd frontend && npm run dev

test: ## Ejecuta todos los tests (backend + frontend)
	@echo "🧪 Ejecutando tests del backend..."
	@cd backend && go test ./...
	@echo ""
	@echo "🧪 Ejecutando tests del frontend..."
	@cd frontend && npm run test:unit

lint: ## Ejecuta linters en todo el proyecto
	@echo "🔍 Ejecutando linters del backend..."
	@cd backend && go fmt ./...
	@echo ""
	@echo "🔍 Ejecutando linters del frontend..."
	@cd frontend && npm run lint

build: ## [Opcional] Compila el proyecto para producción
	@echo "🔨 Compilando backend..."
	@cd backend && go build -o ../bin/softpharos main.go
	@echo ""
	@echo "🔨 Compilando frontend..."
	@cd frontend && npm run build
	@echo ""
	@echo "✅ Build completado:"
	@echo "  • Backend:  bin/softpharos"
	@echo "  • Frontend: frontend/dist/"

db-reset: ## Reinicia la base de datos desde cero
	@bash scripts/re_init_bd.sh

clean: ## Limpia archivos generados y temporales
	@echo "🧹 Limpiando archivos temporales..."
	@cd backend && go clean
	@rm -rf bin/
	@rm -rf frontend/dist/
	@rm -rf frontend/node_modules/
	@echo "✅ Limpieza completada"

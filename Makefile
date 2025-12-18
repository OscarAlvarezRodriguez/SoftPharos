.PHONY: help setup dev dev-backend dev-frontend test clean db-reset

help: ## Muestra esta ayuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Configuración inicial completa
	bash scripts/setup.sh

dev-backend: ## Inicia el backend en modo desarrollo
	cd backend && go run main.go

dev-frontend: ## Inicia el frontend en modo desarrollo
	cd frontend && npm run dev

dev: ## Inicia backend y frontend (requiere 2 terminales)
	@echo "⚠️  Ejecuta en terminales separadas:"
	@echo "   Terminal 1: make dev-backend"
	@echo "   Terminal 2: make dev-frontend"

test-backend: ## Ejecuta tests del backend
	cd backend && go test ./...

test-frontend: ## Ejecuta tests del frontend
	cd frontend && npm run test:unit

test: test-backend test-frontend ## Ejecuta todos los tests

db-reset: ## Reinicia la base de datos
	bash scripts/re_init_bd.sh

clean: ## Limpia archivos temporales
	cd backend && go clean
	cd frontend && rm -rf dist node_modules/.vite

lint-backend: ## Ejecuta linter en backend
	cd backend && go fmt ./...

lint-frontend: ## Ejecuta linter en frontend
	cd frontend && npm run lint

build-backend: ## Compila el backend
	cd backend && go build -o ../bin/softpharos main.go

build-frontend: ## Compila el frontend para producción
	cd frontend && npm run build

build: build-backend build-frontend ## Compila todo el proyecto

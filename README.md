# 🌟 SoftPharos

> Plataforma educativa para visualizar y documentar el proceso de desarrollo de proyectos de software

**Universidad Nacional de Colombia - Ingeniería de Software**

## 👥 Equipo

| Nombre | Correo |
|--------|--------|
| Oscar Eduardo Alvarez Rodriguez | oalvarezr@unal.edu.co |
| Silvana Suarez Carvajal | sisuarezc@unal.edu.co |

## 📝 Descripción

SoftPharos permite a los estudiantes documentar e interactuar con el desarrollo de sus proyectos de software, mostrando que el desarrollo es un **proceso iterativo** donde cada decisión forma parte del aprendizaje.

## 🏗️ Arquitectura

- **Backend**: API REST en Go (puerto 8080) con arquitectura hexagonal
- **Frontend**: Vue 3 + Vite (puerto 5173)
- **Base de datos**: PostgreSQL en Docker

## 🚀 Inicio rápido

### Prerrequisitos

- Docker y Docker Compose
- Go 1.24+
- Node.js 20+
- Archivo `.env` en la raíz (solicitar al equipo)

### Instalación

```bash
# Clonar repositorio
git clone <url-repositorio>
cd SoftPharos

# Ejecutar setup completo
bash scripts/setup.sh
```

### Desarrollo

Abre **dos terminales**:

```bash
# Terminal 1 - Backend
make dev-backend

# Terminal 2 - Frontend
make dev-frontend
```

**URLs:**
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

> 💡 Tip: Ejecuta `make dev` para ver estas instrucciones

## 📁 Estructura

```
SoftPharos/
├── backend/          # API en Go
├── frontend/         # Aplicación Vue 3
├── docs/             # Documentación del proyecto
├── scripts/          # Scripts de utilidad
└── docker-compose.yml
```

## 🛠️ Comandos útiles

```bash
make help           # Ver todos los comandos disponibles
make dev            # Ver instrucciones de desarrollo
make dev-backend    # Iniciar backend (Terminal 1)
make dev-frontend   # Iniciar frontend (Terminal 2)
make test           # Ejecutar todos los tests
make lint           # Ejecutar linters
make db-reset       # Reiniciar base de datos
make clean          # Limpiar archivos temporales
```

> **Nota:** `make build` existe pero es opcional, solo para compilar producción

## 📚 Documentación adicional

- [Guía de inicio rápido](docs/INICIO_RAPIDO.md) - Para nuevos desarrolladores
- [Backend README](backend/README.md) - Arquitectura y API
- [Frontend README](frontend/README.md) - Componentes y vistas
- [Diagramas técnicos](docs/Diagramas/) - Arquitectura y BD

# SoftPharos Backend API

Backend del proyecto SoftPharos desarrollado en Go con arquitectura hexagonal (Clean Architecture).

## 🏗️ Arquitectura

El proyecto sigue los principios de Clean Architecture con la siguiente estructura:

```
Proyecto/Backend/
├── cmd/
│   ├── app/              # Aplicación principal
│   │   ├── main.go       # Punto de entrada
│   │   └── urlMapping.go # Configuración de rutas
│   ├── buildingAPI/      # Inyección de dependencias
│   │   └── container.go  # Contenedor de dependencias
│   └── bd/               # Scripts de base de datos
├── internal/
│   ├── controllers/      # Capa de presentación (HTTP handlers)
│   │   └── project/      # Controladores de proyectos
│   ├── core/
│   │   ├── domain/       # Entidades de dominio
│   │   ├── ports/        # Interfaces (contratos)
│   │   │   ├── repository/  # Interfaces de repositorios
│   │   │   └── services/    # Interfaces de servicios
│   │   ├── repository/   # Implementaciones de repositorios
│   │   └── services/     # Lógica de negocio
│   └── infra/
│       └── databases/    # Infraestructura de persistencia
│           ├── mappers/  # Mapeo entre modelos y dominio
│           └── models/   # Modelos de GORM
```

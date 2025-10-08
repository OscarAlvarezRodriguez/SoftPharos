# Backend - SoftPharos

Backend del proyecto SoftPharos implementado en Go con arquitectura hexagonal.

## 📁 Estructura del Proyecto

```
Backend/
├── main.go                          # Punto de entrada de la aplicación
├── go.mod                           # Dependencias del proyecto
├── internal/
│   ├── core/
│   │   └── domain/                  # Entidades de dominio
│   │       ├── role.go
│   │       ├── user.go
│   │       ├── project.go
│   │       ├── project_member.go
│   │       ├── milestone.go
│   │       ├── deliverable.go
│   │       ├── feedback.go
│   │       ├── comment.go
│   │       └── reaction.go
│   └── infra/
│       └── postgres/                # Infraestructura de base de datos
│           ├── client.go            # Cliente PostgreSQL con GORM
│           └── role_repository.go   # Repositorio de ejemplo
└── cmd/
    └── bd/
        └── init.sql                 # Script de inicialización de BD
```

## 🛠️ Tecnologías

- **Go 1.24.3**
- **GORM** - ORM para PostgreSQL
- **PostgreSQL** - Base de datos relacional
- **godotenv** - Gestión de variables de entorno

## 🚀 Configuración

### Variables de Entorno

Crear un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=tu_usuario
DB_PASSWORD=tu_contraseña
DB_NAME=softpharos
DB_SSLMODE=disable
```

### Instalación de Dependencias

```bash
go mod tidy
```

### Ejecutar el Proyecto

```bash
go run main.go
```

### Compilar

```bash
go build -o bin/main main.go
```

## 🏗️ Arquitectura Hexagonal

El proyecto sigue los principios de arquitectura hexagonal:

### Core (Dominio)
- **`internal/core/domain/`**: Contiene las entidades de negocio sin dependencias externas
- Cada entidad tiene tags GORM para mapeo con la base de datos
- Las relaciones entre entidades están definidas con foreign keys

### Infraestructura
- **`internal/infra/postgres/`**: Implementación concreta del acceso a datos
- **`client.go`**: Gestión de conexión con PostgreSQL usando GORM
- **`*_repository.go`**: Repositorios que implementan operaciones CRUD

## 📊 Entidades del Dominio

### Principales

1. **Role**: Roles de usuarios en el sistema
2. **User**: Usuarios del sistema
3. **Project**: Proyectos
4. **ProjectMember**: Miembros de un proyecto
5. **Milestone**: Hitos de un proyecto
6. **Deliverable**: Entregables de un hito
7. **Feedback**: Retroalimentación de profesores
8. **Comment**: Comentarios en hitos
9. **Reaction**: Reacciones a hitos

## 🔧 Uso del ORM (GORM)

### Ejemplo Básico

```go
// Crear cliente
client, err := postgres.NewClientFromEnv()
if err != nil {
    log.Fatal(err)
}
defer client.Close()

// Obtener registros
var users []domain.User
client.DB.Find(&users)

// Obtener con relaciones (Preload)
client.DB.Preload("Role").Find(&users)

// Crear un registro
newUser := domain.User{
    Email: "user@example.com",
    Password: "hashed_password",
    RoleID: 1,
}
client.DB.Create(&newUser)
```

### Repositorios

Los repositorios encapsulan las operaciones de acceso a datos:

```go
roleRepo := postgres.NewRoleRepository(client)

// Obtener todos los roles
roles, err := roleRepo.GetAll(context.Background())

// Obtener por ID
role, err := roleRepo.GetByID(context.Background(), 1)

// Crear
err := roleRepo.Create(context.Background(), &newRole)
```

## ✨ Ventajas de GORM

- ✅ **No SQL quemado**: Todas las queries se generan automáticamente
- ✅ **Type-safe**: Tipado fuerte de Go
- ✅ **Migraciones**: Soporte para migraciones automáticas
- ✅ **Relaciones**: Manejo automático de foreign keys y preload
- ✅ **Hooks**: Callbacks antes/después de operaciones
- ✅ **Connection pooling**: Gestión automática de conexiones

## 📝 Próximos Pasos

- [ ] Implementar repositorios para todas las entidades
- [ ] Agregar capa de servicios (use cases)
- [ ] Implementar API REST con Gin o Echo
- [ ] Agregar tests unitarios
- [ ] Implementar autenticación JWT
- [ ] Agregar validaciones de datos
- [ ] Configurar migraciones automáticas

## 🤝 Contribución

Este es un proyecto académico de SoftPharos.

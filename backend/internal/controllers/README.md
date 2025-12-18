# Ejemplos de Respuestas Estandarizadas de la API

Todas las respuestas de la API siguen un formato estandarizado con la siguiente estructura:

```json
{
  "success": boolean,
  "data": object | null,
  "error": {
    "code": string,
    "message": string
  } | null,
  "timestamp": string (RFC3339)
}
```

## 📗 Respuestas Exitosas (Success)

### GET - Obtener lista de recursos (200 OK)
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Proyecto A",
      "objective": "Desarrollo de API",
      "created_by": 1,
      "created_at": "2024-11-07T10:00:00Z",
      "updated_at": "2024-11-07T10:00:00Z"
    },
    {
      "id": 2,
      "name": "Proyecto B",
      "objective": "Sistema de gestión",
      "created_by": 2,
      "created_at": "2024-11-07T11:00:00Z",
      "updated_at": "2024-11-07T11:00:00Z"
    }
  ],
  "timestamp": "2024-11-07T12:30:45Z"
}
```

### GET - Obtener un recurso específico (200 OK)
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Admin",
    "description": "Administrador del sistema",
    "created_at": "2024-11-07T10:00:00Z"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

### POST - Crear recurso (201 Created)
```json
{
  "success": true,
  "data": {
    "id": 3,
    "name": "Nuevo Proyecto",
    "objective": "Implementar microservicios",
    "created_by": 1,
    "created_at": "2024-11-07T12:30:45Z",
    "updated_at": "2024-11-07T12:30:45Z"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

### PUT - Actualizar recurso (200 OK)
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Proyecto Actualizado",
    "objective": "Nueva descripción",
    "created_by": 1,
    "created_at": "2024-11-07T10:00:00Z",
    "updated_at": "2024-11-07T12:30:45Z"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

### DELETE - Eliminar recurso (200 OK)
```json
{
  "success": true,
  "data": {
    "message": "Proyecto eliminado exitosamente"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

## 📕 Respuestas de Error

### 400 Bad Request - Solicitud inválida
```json
{
  "success": false,
  "error": {
    "code": "BAD_REQUEST",
    "message": "Key: 'CreateProjectRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

### 400 Bad Request - ID inválido
```json
{
  "success": false,
  "error": {
    "code": "INVALID_ID",
    "message": "El ID debe ser un número válido"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

### 404 Not Found - Recurso no encontrado
```json
{
  "success": false,
  "error": {
    "code": "NOT_FOUND",
    "message": "Proyecto no encontrado"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

### 500 Internal Server Error - Error interno
```json
{
  "success": false,
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "database connection failed"
  },
  "timestamp": "2024-11-07T12:30:45Z"
}
```

## 🔧 Uso del ResponseBuilder en Controllers

```go
// Respuesta exitosa
controllers.Response.Success(ctx, http.StatusOK, data)

// Errores comunes
controllers.Response.BadRequest(ctx, "Mensaje de error")
controllers.Response.NotFound(ctx, "Recurso no encontrado")
controllers.Response.InternalError(ctx, err.Error())
controllers.Response.InvalidID(ctx, "ID inválido")

// Error personalizado
controllers.Response.Error(ctx, statusCode, "CUSTOM_CODE", "Mensaje")
```

## 📋 Códigos de Error Disponibles

| Código | HTTP Status | Descripción |
|--------|-------------|-------------|
| `INTERNAL_ERROR` | 500 | Error interno del servidor |
| `NOT_FOUND` | 404 | Recurso no encontrado |
| `BAD_REQUEST` | 400 | Solicitud malformada |
| `INVALID_ID` | 400 | ID con formato inválido |
| `INVALID_REQUEST` | 400 | Datos de request inválidos |
| `UNAUTHORIZED` | 401 | No autorizado |
| `FORBIDDEN` | 403 | Acceso prohibido |

## ✅ Ventajas de la Estandarización

1. **Consistencia**: Todas las respuestas tienen la misma estructura
2. **Predictibilidad**: El frontend sabe exactamente qué esperar
3. **Facilidad de debug**: El campo `timestamp` ayuda en la trazabilidad
4. **Manejo de errores simplificado**: Estructura única para todos los errores
5. **Type-safe**: Fácil de tipar en TypeScript/JavaScript en el frontend

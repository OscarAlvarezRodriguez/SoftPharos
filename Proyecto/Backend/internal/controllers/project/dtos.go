package project

import "time"

// CreateProjectRequest representa la petición para crear un proyecto
type CreateProjectRequest struct {
	Name      *string `json:"name" binding:"required"`
	Objective *string `json:"objective"`
	CreatedBy int     `json:"created_by" binding:"required"`
}

// UpdateProjectRequest representa la petición para actualizar un proyecto
type UpdateProjectRequest struct {
	Name      *string `json:"name"`
	Objective *string `json:"objective"`
}

// ProjectResponse representa la respuesta de un proyecto
type ProjectResponse struct {
	ID        int              `json:"id"`
	Name      *string          `json:"name"`
	Objective *string          `json:"objective"`
	CreatedBy int              `json:"created_by"`
	Creator   *CreatorResponse `json:"creator,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// CreatorResponse representa la información básica del creador
type CreatorResponse struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

// ErrorResponse representa una respuesta de error
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// SuccessResponse representa una respuesta exitosa genérica
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

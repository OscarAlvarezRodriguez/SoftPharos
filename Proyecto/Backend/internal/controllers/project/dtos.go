package project

import "time"

type CreateProjectRequest struct {
	Name      *string `json:"name" binding:"required"`
	Objective *string `json:"objective"`
	CreatedBy int     `json:"created_by" binding:"required"`
}

type UpdateProjectRequest struct {
	Name      *string `json:"name"`
	Objective *string `json:"objective"`
}

type ProjectResponse struct {
	ID        int            `json:"id"`
	Name      *string        `json:"name"`
	Objective *string        `json:"objective"`
	CreatedBy int            `json:"created_by"`
	Owner     *OwnerResponse `json:"owner,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type OwnerResponse struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

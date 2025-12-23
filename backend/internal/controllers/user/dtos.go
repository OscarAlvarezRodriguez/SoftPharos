package user

import "time"

type CreateUserRequest struct {
	Name       *string `json:"name"`
	Email      string  `json:"email" binding:"required,email"`
	ProviderID string  `json:"provider_id" binding:"required"` // Google sub
	RoleID     int     `json:"role_id" binding:"required"`
	PictureURL *string `json:"picture_url,omitempty"`
}

type UpdateUserRequest struct {
	Name   *string `json:"name"`
	RoleID *int    `json:"role_id"`
}

type UserResponse struct {
	ID         int           `json:"id"`
	Name       *string       `json:"name"`
	Email      string        `json:"email"`
	RoleID     int           `json:"role_id"`
	Role       *RoleResponse `json:"role,omitempty"`
	PictureURL *string       `json:"picture_url,omitempty"`
	CreatedAt  time.Time     `json:"created_at"`
}

type RoleResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

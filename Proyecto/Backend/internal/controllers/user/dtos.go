package user

import "time"

type CreateUserRequest struct {
	Name     *string `json:"name"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=6"`
	RoleID   int     `json:"role_id" binding:"required"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Password *string `json:"password,omitempty"`
	RoleID   *int    `json:"role_id"`
}

type UserResponse struct {
	ID        int           `json:"id"`
	Name      *string       `json:"name"`
	Email     string        `json:"email"`
	RoleID    int           `json:"role_id"`
	Role      *RoleResponse `json:"role,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
}

type RoleResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

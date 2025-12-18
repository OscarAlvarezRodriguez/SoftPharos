package project_member

import "time"

type CreateProjectMemberRequest struct {
	ProjectID int     `json:"project_id" binding:"required"`
	UserID    int     `json:"user_id" binding:"required"`
	Role      *string `json:"role"`
}

type UpdateProjectMemberRequest struct {
	Role *string `json:"role"`
}

type ProjectMemberResponse struct {
	ID        int              `json:"id"`
	ProjectID int              `json:"project_id"`
	Project   *ProjectResponse `json:"project,omitempty"`
	UserID    int              `json:"user_id"`
	User      *UserResponse    `json:"user,omitempty"`
	Role      *string          `json:"role"`
	JoinedAt  time.Time        `json:"joined_at"`
}

type ProjectResponse struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

type UserResponse struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

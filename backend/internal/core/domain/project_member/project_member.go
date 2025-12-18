package project_member

import (
	"time"

	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/domain/user"
)

type ProjectMember struct {
	ID        int
	ProjectID int
	Project   *project.Project
	UserID    int
	User      *user.User
	Role      *string
	JoinedAt  time.Time
}

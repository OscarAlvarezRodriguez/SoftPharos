package user

import (
	"time"

	"softpharos/internal/core/domain/role"
)

type User struct {
	ID         int
	Name       *string
	Email      string
	ProviderID string
	RoleID     int
	Role       *role.Role
	PictureURL *string
	CreatedAt  time.Time
}

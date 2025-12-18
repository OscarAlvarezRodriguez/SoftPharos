package user

import (
	"time"

	"softpharos/internal/core/domain/role"
)

type User struct {
	ID        int
	Name      *string
	Email     string
	Password  string
	RoleID    int
	Role      *role.Role
	CreatedAt time.Time
}

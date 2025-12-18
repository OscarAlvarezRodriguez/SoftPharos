package project

import (
	"time"

	"softpharos/internal/core/domain/user"
)

type Project struct {
	ID        int
	Name      *string
	Objective *string
	CreatedBy int
	Owner     *user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}

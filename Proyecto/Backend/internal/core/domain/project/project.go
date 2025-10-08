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
	Creator   *user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}

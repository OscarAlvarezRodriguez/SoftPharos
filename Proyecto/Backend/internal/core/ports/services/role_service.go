package services

import (
	"context"
	"softpharos/internal/core/domain/role"
)

type RoleService interface {
	GetAllRoles(ctx context.Context) ([]role.Role, error)
	GetRoleByID(ctx context.Context, id int) (*role.Role, error)
	GetRoleByName(ctx context.Context, name string) (*role.Role, error)
}

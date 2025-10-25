package repository

import (
	"context"
	"softpharos/internal/core/domain/role"
)

// RoleRepository define el contrato para las operaciones de persistencia de roles
type RoleRepository interface {
	GetAll(ctx context.Context) ([]role.Role, error)
	GetByID(ctx context.Context, id int) (*role.Role, error)
	GetByName(ctx context.Context, name string) (*role.Role, error)
	Create(ctx context.Context, role *role.Role) error
	Update(ctx context.Context, role *role.Role) error
	Delete(ctx context.Context, id int) error
}

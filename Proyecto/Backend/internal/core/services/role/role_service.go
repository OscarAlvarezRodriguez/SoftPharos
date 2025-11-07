package role

import (
	"context"
	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	roleRepo repository.RoleRepository
}

func New(roleRepo repository.RoleRepository) services.RoleService {
	return &Service{
		roleRepo: roleRepo,
	}
}

func (s *Service) GetAllRoles(ctx context.Context) ([]role.Role, error) {
	return s.roleRepo.GetAll(ctx)
}

func (s *Service) GetRoleByID(ctx context.Context, id int) (*role.Role, error) {
	return s.roleRepo.GetByID(ctx, id)
}

func (s *Service) GetRoleByName(ctx context.Context, name string) (*role.Role, error) {
	return s.roleRepo.GetByName(ctx, name)
}

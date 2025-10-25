package repository

import (
	"context"
	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type roleRepository struct {
	client *databases.Client
}

func NewRoleRepository(client *databases.Client) repository.RoleRepository {
	return &roleRepository{client: client}
}

func (r *roleRepository) GetAll(ctx context.Context) ([]role.Role, error) {
	var roleModels []models.RoleModel
	result := r.client.DB.WithContext(ctx).Find(&roleModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.RoleListToDomain(roleModels), nil
}

func (r *roleRepository) GetByID(ctx context.Context, id int) (*role.Role, error) {
	var roleModel models.RoleModel
	result := r.client.DB.WithContext(ctx).First(&roleModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.RoleToDomain(&roleModel), nil
}

func (r *roleRepository) GetByName(ctx context.Context, name string) (*role.Role, error) {
	var roleModel models.RoleModel
	result := r.client.DB.WithContext(ctx).Where("name = ?", name).First(&roleModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.RoleToDomain(&roleModel), nil
}

func (r *roleRepository) Create(ctx context.Context, domainRole *role.Role) error {
	roleModel := mappers.RoleToModel(domainRole)
	result := r.client.DB.WithContext(ctx).Create(roleModel)
	if result.Error != nil {
		return result.Error
	}

	domainRole.ID = roleModel.ID
	return nil
}

func (r *roleRepository) Update(ctx context.Context, domainRole *role.Role) error {
	roleModel := mappers.RoleToModel(domainRole)
	return r.client.DB.WithContext(ctx).Save(roleModel).Error
}

func (r *roleRepository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.RoleModel{}, id).Error
}

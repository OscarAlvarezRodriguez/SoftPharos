package role

import (
	"context"
	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.RoleRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]role.Role, error) {
	var roleModels []models.RoleModel
	result := r.client.DB.WithContext(ctx).Find(&roleModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.RoleListToDomain(roleModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*role.Role, error) {
	var roleModel models.RoleModel
	result := r.client.DB.WithContext(ctx).First(&roleModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.RoleToDomain(&roleModel), nil
}

func (r *Repository) GetByName(ctx context.Context, name string) (*role.Role, error) {
	var roleModel models.RoleModel
	result := r.client.DB.WithContext(ctx).Where("name = ?", name).First(&roleModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.RoleToDomain(&roleModel), nil
}

func (r *Repository) Create(ctx context.Context, domainRole *role.Role) error {
	roleModel := mappers.RoleToModel(domainRole)
	result := r.client.DB.WithContext(ctx).Create(roleModel)
	if result.Error != nil {
		return result.Error
	}

	domainRole.ID = roleModel.ID
	return nil
}

func (r *Repository) Update(ctx context.Context, domainRole *role.Role) error {
	roleModel := mappers.RoleToModel(domainRole)
	return r.client.DB.WithContext(ctx).Save(roleModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.RoleModel{}, id).Error
}

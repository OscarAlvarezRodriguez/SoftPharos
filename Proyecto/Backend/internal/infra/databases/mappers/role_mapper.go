package mappers

import (
	"softpharos/internal/core/domain/role"
	"softpharos/internal/infra/databases/models"
)

func RoleToDomain(model *models.RoleModel) *role.Role {
	if model == nil {
		return nil
	}

	return &role.Role{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
	}
}

func RoleToModel(domain *role.Role) *models.RoleModel {
	if domain == nil {
		return nil
	}

	return &models.RoleModel{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
	}
}

func RoleListToDomain(modelList []models.RoleModel) []role.Role {
	domainList := make([]role.Role, len(modelList))
	for i, model := range modelList {
		domainList[i] = *RoleToDomain(&model)
	}
	return domainList
}

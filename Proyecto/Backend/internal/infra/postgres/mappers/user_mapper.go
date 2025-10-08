package mappers

import (
	"softpharos/internal/core/domain/user"
	"softpharos/internal/infra/postgres/models"
)

// UserToDomain convierte un modelo de persistencia a entidad de dominio
func UserToDomain(model *models.UserModel) *user.User {
	if model == nil {
		return nil
	}

	return &user.User{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		RoleID:    model.RoleID,
		Role:      RoleToDomain(model.Role),
		CreatedAt: model.CreatedAt,
	}
}

// UserToModel convierte una entidad de dominio a modelo de persistencia
func UserToModel(domain *user.User) *models.UserModel {
	if domain == nil {
		return nil
	}

	return &models.UserModel{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		RoleID:    domain.RoleID,
		Role:      RoleToModel(domain.Role),
		CreatedAt: domain.CreatedAt,
	}
}

// UserListToDomain convierte una lista de modelos a lista de entidades
func UserListToDomain(modelList []models.UserModel) []user.User {
	domainList := make([]user.User, len(modelList))
	for i, model := range modelList {
		domainList[i] = *UserToDomain(&model)
	}
	return domainList
}

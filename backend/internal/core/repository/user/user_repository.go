package user

import (
	"context"
	"softpharos/internal/core/domain/user"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.UserRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]user.User, error) {
	var userModels []models.UserModel
	result := r.client.DB.WithContext(ctx).Preload("Role").Find(&userModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.UserListToDomain(userModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*user.User, error) {
	var userModel models.UserModel
	result := r.client.DB.WithContext(ctx).Preload("Role").First(&userModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.UserToDomain(&userModel), nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var userModel models.UserModel
	result := r.client.DB.WithContext(ctx).Preload("Role").Where("email = ?", email).First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.UserToDomain(&userModel), nil
}

func (r *Repository) Create(ctx context.Context, domainUser *user.User) error {
	userModel := mappers.UserToModel(domainUser)
	result := r.client.DB.WithContext(ctx).Create(userModel)
	if result.Error != nil {
		return result.Error
	}

	domainUser.ID = userModel.ID
	domainUser.CreatedAt = userModel.CreatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainUser *user.User) error {
	userModel := mappers.UserToModel(domainUser)
	return r.client.DB.WithContext(ctx).Save(userModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.UserModel{}, id).Error
}

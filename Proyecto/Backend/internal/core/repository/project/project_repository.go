package project

import (
	"context"
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.ProjectRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]project.Project, error) {
	var projectModels []models.ProjectModel
	result := r.client.DB.WithContext(ctx).Preload("Owner").Find(&projectModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ProjectListToDomain(projectModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*project.Project, error) {
	var projectModel models.ProjectModel
	result := r.client.DB.WithContext(ctx).Preload("Owner").First(&projectModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ProjectToDomain(&projectModel), nil
}

func (r *Repository) GetByOwner(ctx context.Context, ownerID int) ([]project.Project, error) {
	var projectModels []models.ProjectModel
	result := r.client.DB.WithContext(ctx).
		Preload("Owner").
		Where("created_by = ?", ownerID).
		Find(&projectModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ProjectListToDomain(projectModels), nil
}

func (r *Repository) Create(ctx context.Context, domainProject *project.Project) error {
	projectModel := mappers.ProjectToModel(domainProject)
	result := r.client.DB.WithContext(ctx).Create(projectModel)
	if result.Error != nil {
		return result.Error
	}

	domainProject.ID = projectModel.ID
	domainProject.CreatedAt = projectModel.CreatedAt
	domainProject.UpdatedAt = projectModel.UpdatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainProject *project.Project) error {
	projectModel := mappers.ProjectToModel(domainProject)
	return r.client.DB.WithContext(ctx).Save(projectModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.ProjectModel{}, id).Error
}

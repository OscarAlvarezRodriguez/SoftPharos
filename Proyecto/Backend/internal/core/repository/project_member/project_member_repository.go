package project_member

import (
	"context"
	"softpharos/internal/core/domain/project_member"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.ProjectMemberRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]project_member.ProjectMember, error) {
	var projectMemberModels []models.ProjectMemberModel
	result := r.client.DB.WithContext(ctx).Preload("Project").Preload("User").Find(&projectMemberModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ProjectMemberListToDomain(projectMemberModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*project_member.ProjectMember, error) {
	var projectMemberModel models.ProjectMemberModel
	result := r.client.DB.WithContext(ctx).Preload("Project").Preload("User").First(&projectMemberModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ProjectMemberToDomain(&projectMemberModel), nil
}

func (r *Repository) GetByProjectID(ctx context.Context, projectID int) ([]project_member.ProjectMember, error) {
	var projectMemberModels []models.ProjectMemberModel
	result := r.client.DB.WithContext(ctx).
		Preload("Project").
		Preload("User").
		Where("project_id = ?", projectID).
		Find(&projectMemberModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ProjectMemberListToDomain(projectMemberModels), nil
}

func (r *Repository) Create(ctx context.Context, domainProjectMember *project_member.ProjectMember) error {
	projectMemberModel := mappers.ProjectMemberToModel(domainProjectMember)
	result := r.client.DB.WithContext(ctx).Create(projectMemberModel)
	if result.Error != nil {
		return result.Error
	}

	domainProjectMember.ID = projectMemberModel.ID
	domainProjectMember.JoinedAt = projectMemberModel.JoinedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainProjectMember *project_member.ProjectMember) error {
	projectMemberModel := mappers.ProjectMemberToModel(domainProjectMember)
	return r.client.DB.WithContext(ctx).Save(projectMemberModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.ProjectMemberModel{}, id).Error
}

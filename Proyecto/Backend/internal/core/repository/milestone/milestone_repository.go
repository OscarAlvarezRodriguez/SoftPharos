package milestone

import (
	"context"
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.MilestoneRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]milestone.Milestone, error) {
	var milestoneModels []models.MilestoneModel
	result := r.client.DB.WithContext(ctx).Preload("Project").Find(&milestoneModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.MilestoneListToDomain(milestoneModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*milestone.Milestone, error) {
	var milestoneModel models.MilestoneModel
	result := r.client.DB.WithContext(ctx).Preload("Project").First(&milestoneModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.MilestoneToDomain(&milestoneModel), nil
}

func (r *Repository) GetByProjectID(ctx context.Context, projectID int) ([]milestone.Milestone, error) {
	var milestoneModels []models.MilestoneModel
	result := r.client.DB.WithContext(ctx).
		Preload("Project").
		Where("project_id = ?", projectID).
		Find(&milestoneModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.MilestoneListToDomain(milestoneModels), nil
}

func (r *Repository) Create(ctx context.Context, domainMilestone *milestone.Milestone) error {
	milestoneModel := mappers.MilestoneToModel(domainMilestone)
	result := r.client.DB.WithContext(ctx).Create(milestoneModel)
	if result.Error != nil {
		return result.Error
	}

	domainMilestone.ID = milestoneModel.ID
	domainMilestone.CreatedAt = milestoneModel.CreatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainMilestone *milestone.Milestone) error {
	milestoneModel := mappers.MilestoneToModel(domainMilestone)
	return r.client.DB.WithContext(ctx).Save(milestoneModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.MilestoneModel{}, id).Error
}

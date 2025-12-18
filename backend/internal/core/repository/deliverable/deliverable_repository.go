package deliverable

import (
	"context"
	"softpharos/internal/core/domain/deliverable"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.DeliverableRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]deliverable.Deliverable, error) {
	var deliverableModels []models.DeliverableModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").Find(&deliverableModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.DeliverableListToDomain(deliverableModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*deliverable.Deliverable, error) {
	var deliverableModel models.DeliverableModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").First(&deliverableModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.DeliverableToDomain(&deliverableModel), nil
}

func (r *Repository) GetByMilestoneID(ctx context.Context, milestoneID int) ([]deliverable.Deliverable, error) {
	var deliverableModels []models.DeliverableModel
	result := r.client.DB.WithContext(ctx).
		Preload("Milestone").
		Where("milestone_id = ?", milestoneID).
		Find(&deliverableModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.DeliverableListToDomain(deliverableModels), nil
}

func (r *Repository) Create(ctx context.Context, domainDeliverable *deliverable.Deliverable) error {
	deliverableModel := mappers.DeliverableToModel(domainDeliverable)
	result := r.client.DB.WithContext(ctx).Create(deliverableModel)
	if result.Error != nil {
		return result.Error
	}

	domainDeliverable.ID = deliverableModel.ID
	domainDeliverable.CreatedAt = deliverableModel.CreatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainDeliverable *deliverable.Deliverable) error {
	deliverableModel := mappers.DeliverableToModel(domainDeliverable)
	return r.client.DB.WithContext(ctx).Save(deliverableModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.DeliverableModel{}, id).Error
}

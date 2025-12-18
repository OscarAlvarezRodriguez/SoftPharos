package reaction

import (
	"context"
	"softpharos/internal/core/domain/reaction"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.ReactionRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]reaction.Reaction, error) {
	var reactionModels []models.ReactionModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").Preload("User").Find(&reactionModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ReactionListToDomain(reactionModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*reaction.Reaction, error) {
	var reactionModel models.ReactionModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").Preload("User").First(&reactionModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ReactionToDomain(&reactionModel), nil
}

func (r *Repository) GetByMilestoneID(ctx context.Context, milestoneID int) ([]reaction.Reaction, error) {
	var reactionModels []models.ReactionModel
	result := r.client.DB.WithContext(ctx).
		Preload("Milestone").
		Preload("User").
		Where("milestone_id = ?", milestoneID).
		Find(&reactionModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.ReactionListToDomain(reactionModels), nil
}

func (r *Repository) Create(ctx context.Context, domainReaction *reaction.Reaction) error {
	reactionModel := mappers.ReactionToModel(domainReaction)
	result := r.client.DB.WithContext(ctx).Create(reactionModel)
	if result.Error != nil {
		return result.Error
	}

	domainReaction.ID = reactionModel.ID
	domainReaction.CreatedAt = reactionModel.CreatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainReaction *reaction.Reaction) error {
	reactionModel := mappers.ReactionToModel(domainReaction)
	return r.client.DB.WithContext(ctx).Save(reactionModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.ReactionModel{}, id).Error
}

package comment

import (
	"context"
	"softpharos/internal/core/domain/comment"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.CommentRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]comment.Comment, error) {
	var commentModels []models.CommentModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").Preload("User").Find(&commentModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.CommentListToDomain(commentModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*comment.Comment, error) {
	var commentModel models.CommentModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").Preload("User").First(&commentModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.CommentToDomain(&commentModel), nil
}

func (r *Repository) GetByMilestoneID(ctx context.Context, milestoneID int) ([]comment.Comment, error) {
	var commentModels []models.CommentModel
	result := r.client.DB.WithContext(ctx).
		Preload("Milestone").
		Preload("User").
		Where("milestone_id = ?", milestoneID).
		Find(&commentModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.CommentListToDomain(commentModels), nil
}

func (r *Repository) Create(ctx context.Context, domainComment *comment.Comment) error {
	commentModel := mappers.CommentToModel(domainComment)
	result := r.client.DB.WithContext(ctx).Create(commentModel)
	if result.Error != nil {
		return result.Error
	}

	domainComment.ID = commentModel.ID
	domainComment.CreatedAt = commentModel.CreatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainComment *comment.Comment) error {
	commentModel := mappers.CommentToModel(domainComment)
	return r.client.DB.WithContext(ctx).Save(commentModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.CommentModel{}, id).Error
}

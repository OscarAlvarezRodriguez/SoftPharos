package feedback

import (
	"context"
	"softpharos/internal/core/domain/feedback"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/mappers"
	"softpharos/internal/infra/databases/models"
)

type Repository struct {
	client *databases.Client
}

func New(client *databases.Client) repository.FeedbackRepository {
	return &Repository{client: client}
}

func (r *Repository) GetAll(ctx context.Context) ([]feedback.Feedback, error) {
	var feedbackModels []models.FeedbackModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").Preload("Professor").Find(&feedbackModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.FeedbackListToDomain(feedbackModels), nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*feedback.Feedback, error) {
	var feedbackModel models.FeedbackModel
	result := r.client.DB.WithContext(ctx).Preload("Milestone").Preload("Professor").First(&feedbackModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.FeedbackToDomain(&feedbackModel), nil
}

func (r *Repository) GetByMilestoneID(ctx context.Context, milestoneID int) ([]feedback.Feedback, error) {
	var feedbackModels []models.FeedbackModel
	result := r.client.DB.WithContext(ctx).
		Preload("Milestone").
		Preload("Professor").
		Where("milestone_id = ?", milestoneID).
		Find(&feedbackModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.FeedbackListToDomain(feedbackModels), nil
}

func (r *Repository) Create(ctx context.Context, domainFeedback *feedback.Feedback) error {
	feedbackModel := mappers.FeedbackToModel(domainFeedback)
	result := r.client.DB.WithContext(ctx).Create(feedbackModel)
	if result.Error != nil {
		return result.Error
	}

	domainFeedback.ID = feedbackModel.ID
	domainFeedback.CreatedAt = feedbackModel.CreatedAt
	return nil
}

func (r *Repository) Update(ctx context.Context, domainFeedback *feedback.Feedback) error {
	feedbackModel := mappers.FeedbackToModel(domainFeedback)
	return r.client.DB.WithContext(ctx).Save(feedbackModel).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&models.FeedbackModel{}, id).Error
}

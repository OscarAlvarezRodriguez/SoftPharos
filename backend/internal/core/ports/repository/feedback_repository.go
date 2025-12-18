package repository

import (
	"context"
	"softpharos/internal/core/domain/feedback"
)

type FeedbackRepository interface {
	GetAll(ctx context.Context) ([]feedback.Feedback, error)
	GetByID(ctx context.Context, id int) (*feedback.Feedback, error)
	GetByMilestoneID(ctx context.Context, milestoneID int) ([]feedback.Feedback, error)
	Create(ctx context.Context, feedback *feedback.Feedback) error
	Update(ctx context.Context, feedback *feedback.Feedback) error
	Delete(ctx context.Context, id int) error
}

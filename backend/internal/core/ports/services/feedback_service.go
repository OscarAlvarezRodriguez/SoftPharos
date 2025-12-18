package services

import (
	"context"
	"softpharos/internal/core/domain/feedback"
)

type FeedbackService interface {
	GetAllFeedbacks(ctx context.Context) ([]feedback.Feedback, error)
	GetFeedbackByID(ctx context.Context, id int) (*feedback.Feedback, error)
	GetFeedbacksByMilestoneID(ctx context.Context, milestoneID int) ([]feedback.Feedback, error)
	CreateFeedback(ctx context.Context, feedback *feedback.Feedback) error
	UpdateFeedback(ctx context.Context, feedback *feedback.Feedback) error
	DeleteFeedback(ctx context.Context, id int) error
}

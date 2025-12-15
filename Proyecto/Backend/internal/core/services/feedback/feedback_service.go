package feedback

import (
	"context"
	"softpharos/internal/core/domain/feedback"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	feedbackRepo repository.FeedbackRepository
}

func New(feedbackRepo repository.FeedbackRepository) services.FeedbackService {
	return &Service{
		feedbackRepo: feedbackRepo,
	}
}

func (s *Service) GetAllFeedbacks(ctx context.Context) ([]feedback.Feedback, error) {
	return s.feedbackRepo.GetAll(ctx)
}

func (s *Service) GetFeedbackByID(ctx context.Context, id int) (*feedback.Feedback, error) {
	return s.feedbackRepo.GetByID(ctx, id)
}

func (s *Service) GetFeedbacksByMilestoneID(ctx context.Context, milestoneID int) ([]feedback.Feedback, error) {
	return s.feedbackRepo.GetByMilestoneID(ctx, milestoneID)
}

func (s *Service) CreateFeedback(ctx context.Context, f *feedback.Feedback) error {
	return s.feedbackRepo.Create(ctx, f)
}

func (s *Service) UpdateFeedback(ctx context.Context, f *feedback.Feedback) error {
	return s.feedbackRepo.Update(ctx, f)
}

func (s *Service) DeleteFeedback(ctx context.Context, id int) error {
	return s.feedbackRepo.Delete(ctx, id)
}

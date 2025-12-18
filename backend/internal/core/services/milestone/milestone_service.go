package milestone

import (
	"context"
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	milestoneRepo repository.MilestoneRepository
}

func New(milestoneRepo repository.MilestoneRepository) services.MilestoneService {
	return &Service{
		milestoneRepo: milestoneRepo,
	}
}

func (s *Service) GetAllMilestones(ctx context.Context) ([]milestone.Milestone, error) {
	return s.milestoneRepo.GetAll(ctx)
}

func (s *Service) GetMilestoneByID(ctx context.Context, id int) (*milestone.Milestone, error) {
	return s.milestoneRepo.GetByID(ctx, id)
}

func (s *Service) GetMilestonesByProjectID(ctx context.Context, projectID int) ([]milestone.Milestone, error) {
	return s.milestoneRepo.GetByProjectID(ctx, projectID)
}

func (s *Service) CreateMilestone(ctx context.Context, m *milestone.Milestone) error {
	return s.milestoneRepo.Create(ctx, m)
}

func (s *Service) UpdateMilestone(ctx context.Context, m *milestone.Milestone) error {
	return s.milestoneRepo.Update(ctx, m)
}

func (s *Service) DeleteMilestone(ctx context.Context, id int) error {
	return s.milestoneRepo.Delete(ctx, id)
}

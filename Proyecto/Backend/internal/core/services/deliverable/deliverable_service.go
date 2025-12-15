package deliverable

import (
	"context"
	"softpharos/internal/core/domain/deliverable"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	deliverableRepo repository.DeliverableRepository
}

func New(deliverableRepo repository.DeliverableRepository) services.DeliverableService {
	return &Service{
		deliverableRepo: deliverableRepo,
	}
}

func (s *Service) GetAllDeliverables(ctx context.Context) ([]deliverable.Deliverable, error) {
	return s.deliverableRepo.GetAll(ctx)
}

func (s *Service) GetDeliverableByID(ctx context.Context, id int) (*deliverable.Deliverable, error) {
	return s.deliverableRepo.GetByID(ctx, id)
}

func (s *Service) GetDeliverablesByMilestoneID(ctx context.Context, milestoneID int) ([]deliverable.Deliverable, error) {
	return s.deliverableRepo.GetByMilestoneID(ctx, milestoneID)
}

func (s *Service) CreateDeliverable(ctx context.Context, d *deliverable.Deliverable) error {
	return s.deliverableRepo.Create(ctx, d)
}

func (s *Service) UpdateDeliverable(ctx context.Context, d *deliverable.Deliverable) error {
	return s.deliverableRepo.Update(ctx, d)
}

func (s *Service) DeleteDeliverable(ctx context.Context, id int) error {
	return s.deliverableRepo.Delete(ctx, id)
}

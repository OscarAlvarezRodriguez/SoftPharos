package repository

import (
	"context"
	"softpharos/internal/core/domain/deliverable"
)

type DeliverableRepository interface {
	GetAll(ctx context.Context) ([]deliverable.Deliverable, error)
	GetByID(ctx context.Context, id int) (*deliverable.Deliverable, error)
	GetByMilestoneID(ctx context.Context, milestoneID int) ([]deliverable.Deliverable, error)
	Create(ctx context.Context, deliverable *deliverable.Deliverable) error
	Update(ctx context.Context, deliverable *deliverable.Deliverable) error
	Delete(ctx context.Context, id int) error
}

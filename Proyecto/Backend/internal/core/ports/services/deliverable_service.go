package services

import (
	"context"
	"softpharos/internal/core/domain/deliverable"
)

type DeliverableService interface {
	GetAllDeliverables(ctx context.Context) ([]deliverable.Deliverable, error)
	GetDeliverableByID(ctx context.Context, id int) (*deliverable.Deliverable, error)
	GetDeliverablesByMilestoneID(ctx context.Context, milestoneID int) ([]deliverable.Deliverable, error)
	CreateDeliverable(ctx context.Context, deliverable *deliverable.Deliverable) error
	UpdateDeliverable(ctx context.Context, deliverable *deliverable.Deliverable) error
	DeleteDeliverable(ctx context.Context, id int) error
}

package repository

import (
	"context"
	"softpharos/internal/core/domain/milestone"
)

type MilestoneRepository interface {
	GetAll(ctx context.Context) ([]milestone.Milestone, error)
	GetByID(ctx context.Context, id int) (*milestone.Milestone, error)
	GetByProjectID(ctx context.Context, projectID int) ([]milestone.Milestone, error)
	Create(ctx context.Context, milestone *milestone.Milestone) error
	Update(ctx context.Context, milestone *milestone.Milestone) error
	Delete(ctx context.Context, id int) error
}

package services

import (
	"context"
	"softpharos/internal/core/domain/milestone"
)

type MilestoneService interface {
	GetAllMilestones(ctx context.Context) ([]milestone.Milestone, error)
	GetMilestoneByID(ctx context.Context, id int) (*milestone.Milestone, error)
	GetMilestonesByProjectID(ctx context.Context, projectID int) ([]milestone.Milestone, error)
	CreateMilestone(ctx context.Context, milestone *milestone.Milestone) error
	UpdateMilestone(ctx context.Context, milestone *milestone.Milestone) error
	DeleteMilestone(ctx context.Context, id int) error
}

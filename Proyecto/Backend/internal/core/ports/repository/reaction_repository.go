package repository

import (
	"context"
	"softpharos/internal/core/domain/reaction"
)

type ReactionRepository interface {
	GetAll(ctx context.Context) ([]reaction.Reaction, error)
	GetByID(ctx context.Context, id int) (*reaction.Reaction, error)
	GetByMilestoneID(ctx context.Context, milestoneID int) ([]reaction.Reaction, error)
	Create(ctx context.Context, reaction *reaction.Reaction) error
	Update(ctx context.Context, reaction *reaction.Reaction) error
	Delete(ctx context.Context, id int) error
}

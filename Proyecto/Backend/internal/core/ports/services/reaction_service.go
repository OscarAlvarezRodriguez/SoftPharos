package services

import (
	"context"
	"softpharos/internal/core/domain/reaction"
)

type ReactionService interface {
	GetAllReactions(ctx context.Context) ([]reaction.Reaction, error)
	GetReactionByID(ctx context.Context, id int) (*reaction.Reaction, error)
	GetReactionsByMilestoneID(ctx context.Context, milestoneID int) ([]reaction.Reaction, error)
	CreateReaction(ctx context.Context, reaction *reaction.Reaction) error
	UpdateReaction(ctx context.Context, reaction *reaction.Reaction) error
	DeleteReaction(ctx context.Context, id int) error
}

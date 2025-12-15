package reaction

import (
	"context"
	"softpharos/internal/core/domain/reaction"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	reactionRepo repository.ReactionRepository
}

func New(reactionRepo repository.ReactionRepository) services.ReactionService {
	return &Service{
		reactionRepo: reactionRepo,
	}
}

func (s *Service) GetAllReactions(ctx context.Context) ([]reaction.Reaction, error) {
	return s.reactionRepo.GetAll(ctx)
}

func (s *Service) GetReactionByID(ctx context.Context, id int) (*reaction.Reaction, error) {
	return s.reactionRepo.GetByID(ctx, id)
}

func (s *Service) GetReactionsByMilestoneID(ctx context.Context, milestoneID int) ([]reaction.Reaction, error) {
	return s.reactionRepo.GetByMilestoneID(ctx, milestoneID)
}

func (s *Service) CreateReaction(ctx context.Context, r *reaction.Reaction) error {
	return s.reactionRepo.Create(ctx, r)
}

func (s *Service) UpdateReaction(ctx context.Context, r *reaction.Reaction) error {
	return s.reactionRepo.Update(ctx, r)
}

func (s *Service) DeleteReaction(ctx context.Context, id int) error {
	return s.reactionRepo.Delete(ctx, id)
}

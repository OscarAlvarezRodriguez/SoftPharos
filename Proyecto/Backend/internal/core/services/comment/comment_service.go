package comment

import (
	"context"
	"softpharos/internal/core/domain/comment"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	commentRepo repository.CommentRepository
}

func New(commentRepo repository.CommentRepository) services.CommentService {
	return &Service{
		commentRepo: commentRepo,
	}
}

func (s *Service) GetAllComments(ctx context.Context) ([]comment.Comment, error) {
	return s.commentRepo.GetAll(ctx)
}

func (s *Service) GetCommentByID(ctx context.Context, id int) (*comment.Comment, error) {
	return s.commentRepo.GetByID(ctx, id)
}

func (s *Service) GetCommentsByMilestoneID(ctx context.Context, milestoneID int) ([]comment.Comment, error) {
	return s.commentRepo.GetByMilestoneID(ctx, milestoneID)
}

func (s *Service) CreateComment(ctx context.Context, c *comment.Comment) error {
	return s.commentRepo.Create(ctx, c)
}

func (s *Service) UpdateComment(ctx context.Context, c *comment.Comment) error {
	return s.commentRepo.Update(ctx, c)
}

func (s *Service) DeleteComment(ctx context.Context, id int) error {
	return s.commentRepo.Delete(ctx, id)
}

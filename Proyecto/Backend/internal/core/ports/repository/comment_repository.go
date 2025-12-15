package repository

import (
	"context"
	"softpharos/internal/core/domain/comment"
)

type CommentRepository interface {
	GetAll(ctx context.Context) ([]comment.Comment, error)
	GetByID(ctx context.Context, id int) (*comment.Comment, error)
	GetByMilestoneID(ctx context.Context, milestoneID int) ([]comment.Comment, error)
	Create(ctx context.Context, comment *comment.Comment) error
	Update(ctx context.Context, comment *comment.Comment) error
	Delete(ctx context.Context, id int) error
}

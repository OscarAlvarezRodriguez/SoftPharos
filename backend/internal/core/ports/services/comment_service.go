package services

import (
	"context"
	"softpharos/internal/core/domain/comment"
)

type CommentService interface {
	GetAllComments(ctx context.Context) ([]comment.Comment, error)
	GetCommentByID(ctx context.Context, id int) (*comment.Comment, error)
	GetCommentsByMilestoneID(ctx context.Context, milestoneID int) ([]comment.Comment, error)
	CreateComment(ctx context.Context, comment *comment.Comment) error
	UpdateComment(ctx context.Context, comment *comment.Comment) error
	DeleteComment(ctx context.Context, id int) error
}

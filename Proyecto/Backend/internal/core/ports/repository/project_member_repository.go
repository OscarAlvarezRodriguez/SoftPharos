package repository

import (
	"context"
	"softpharos/internal/core/domain/project_member"
)

type ProjectMemberRepository interface {
	GetAll(ctx context.Context) ([]project_member.ProjectMember, error)
	GetByID(ctx context.Context, id int) (*project_member.ProjectMember, error)
	GetByProjectID(ctx context.Context, projectID int) ([]project_member.ProjectMember, error)
	Create(ctx context.Context, projectMember *project_member.ProjectMember) error
	Update(ctx context.Context, projectMember *project_member.ProjectMember) error
	Delete(ctx context.Context, id int) error
}

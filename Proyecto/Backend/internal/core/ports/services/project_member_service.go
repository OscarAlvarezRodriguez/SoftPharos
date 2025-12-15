package services

import (
	"context"
	"softpharos/internal/core/domain/project_member"
)

type ProjectMemberService interface {
	GetAllProjectMembers(ctx context.Context) ([]project_member.ProjectMember, error)
	GetProjectMemberByID(ctx context.Context, id int) (*project_member.ProjectMember, error)
	GetProjectMembersByProjectID(ctx context.Context, projectID int) ([]project_member.ProjectMember, error)
	CreateProjectMember(ctx context.Context, projectMember *project_member.ProjectMember) error
	UpdateProjectMember(ctx context.Context, projectMember *project_member.ProjectMember) error
	DeleteProjectMember(ctx context.Context, id int) error
}

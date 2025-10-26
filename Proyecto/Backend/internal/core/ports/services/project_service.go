package services

import (
	"context"
	"softpharos/internal/core/domain/project"
)

type ProjectService interface {
	GetAllProjects(ctx context.Context) ([]project.Project, error)
	GetProjectByID(ctx context.Context, id int) (*project.Project, error)
	GetProjectsByOwner(ctx context.Context, ownerID int) ([]project.Project, error)
	CreateProject(ctx context.Context, project *project.Project) error
	UpdateProject(ctx context.Context, project *project.Project) error
	DeleteProject(ctx context.Context, id int) error
}

package repository

import (
	"context"
	"softpharos/internal/core/domain/project"
)

// ProjectRepository define el contrato para las operaciones de persistencia de proyectos
type ProjectRepository interface {
	GetAll(ctx context.Context) ([]project.Project, error)
	GetByID(ctx context.Context, id int) (*project.Project, error)
	GetByCreator(ctx context.Context, creatorID int) ([]project.Project, error)
	Create(ctx context.Context, project *project.Project) error
	Update(ctx context.Context, project *project.Project) error
	Delete(ctx context.Context, id int) error
}

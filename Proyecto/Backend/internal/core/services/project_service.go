package services

import (
	"context"
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type projectService struct {
	projectRepo repository.ProjectRepository
}

func New(projectRepo repository.ProjectRepository) services.ProjectService {
	return &projectService{
		projectRepo: projectRepo,
	}
}

func (s *projectService) GetAllProjects(ctx context.Context) ([]project.Project, error) {
	return s.projectRepo.GetAll(ctx)
}

func (s *projectService) GetProjectByID(ctx context.Context, id int) (*project.Project, error) {
	return s.projectRepo.GetByID(ctx, id)
}

func (s *projectService) GetProjectsByOwner(ctx context.Context, ownerID int) ([]project.Project, error) {
	return s.projectRepo.GetByOwner(ctx, ownerID)
}

func (s *projectService) CreateProject(ctx context.Context, proj *project.Project) error {
	return s.projectRepo.Create(ctx, proj)
}

func (s *projectService) UpdateProject(ctx context.Context, proj *project.Project) error {
	return s.projectRepo.Update(ctx, proj)
}

func (s *projectService) DeleteProject(ctx context.Context, id int) error {
	return s.projectRepo.Delete(ctx, id)
}

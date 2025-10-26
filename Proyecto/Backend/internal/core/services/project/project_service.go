package project

import (
	"context"
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	projectRepo repository.ProjectRepository
}

func New(projectRepo repository.ProjectRepository) services.ProjectService {
	return &Service{
		projectRepo: projectRepo,
	}
}

func (s *Service) GetAllProjects(ctx context.Context) ([]project.Project, error) {
	return s.projectRepo.GetAll(ctx)
}

func (s *Service) GetProjectByID(ctx context.Context, id int) (*project.Project, error) {
	return s.projectRepo.GetByID(ctx, id)
}

func (s *Service) GetProjectsByOwner(ctx context.Context, ownerID int) ([]project.Project, error) {
	return s.projectRepo.GetByOwner(ctx, ownerID)
}

func (s *Service) CreateProject(ctx context.Context, proj *project.Project) error {
	return s.projectRepo.Create(ctx, proj)
}

func (s *Service) UpdateProject(ctx context.Context, proj *project.Project) error {
	return s.projectRepo.Update(ctx, proj)
}

func (s *Service) DeleteProject(ctx context.Context, id int) error {
	return s.projectRepo.Delete(ctx, id)
}

package project_member

import (
	"context"
	"softpharos/internal/core/domain/project_member"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	projectMemberRepo repository.ProjectMemberRepository
}

func New(projectMemberRepo repository.ProjectMemberRepository) services.ProjectMemberService {
	return &Service{
		projectMemberRepo: projectMemberRepo,
	}
}

func (s *Service) GetAllProjectMembers(ctx context.Context) ([]project_member.ProjectMember, error) {
	return s.projectMemberRepo.GetAll(ctx)
}

func (s *Service) GetProjectMemberByID(ctx context.Context, id int) (*project_member.ProjectMember, error) {
	return s.projectMemberRepo.GetByID(ctx, id)
}

func (s *Service) GetProjectMembersByProjectID(ctx context.Context, projectID int) ([]project_member.ProjectMember, error) {
	return s.projectMemberRepo.GetByProjectID(ctx, projectID)
}

func (s *Service) CreateProjectMember(ctx context.Context, pm *project_member.ProjectMember) error {
	return s.projectMemberRepo.Create(ctx, pm)
}

func (s *Service) UpdateProjectMember(ctx context.Context, pm *project_member.ProjectMember) error {
	return s.projectMemberRepo.Update(ctx, pm)
}

func (s *Service) DeleteProjectMember(ctx context.Context, id int) error {
	return s.projectMemberRepo.Delete(ctx, id)
}

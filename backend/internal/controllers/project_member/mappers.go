package project_member

import (
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/domain/project_member"
	"softpharos/internal/core/domain/user"
)

func ToProjectMemberDomain(req *CreateProjectMemberRequest) *project_member.ProjectMember {
	return &project_member.ProjectMember{
		ProjectID: req.ProjectID,
		UserID:    req.UserID,
		Role:      req.Role,
	}
}

func ToProjectMemberResponse(pm *project_member.ProjectMember) *ProjectMemberResponse {
	if pm == nil {
		return nil
	}

	response := &ProjectMemberResponse{
		ID:        pm.ID,
		ProjectID: pm.ProjectID,
		UserID:    pm.UserID,
		Role:      pm.Role,
		JoinedAt:  pm.JoinedAt,
	}

	if pm.Project != nil {
		response.Project = ToProjectResponse(pm.Project)
	}

	if pm.User != nil {
		response.User = ToUserResponse(pm.User)
	}

	return response
}

func ToProjectResponse(p *project.Project) *ProjectResponse {
	if p == nil {
		return nil
	}

	return &ProjectResponse{
		ID:   p.ID,
		Name: p.Name,
	}
}

func ToUserResponse(u *user.User) *UserResponse {
	if u == nil {
		return nil
	}

	return &UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func ToProjectMemberListResponse(projectMembers []project_member.ProjectMember) []ProjectMemberResponse {
	responses := make([]ProjectMemberResponse, len(projectMembers))
	for i, pm := range projectMembers {
		responses[i] = *ToProjectMemberResponse(&pm)
	}
	return responses
}

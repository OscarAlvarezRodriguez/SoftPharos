package project

import (
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/domain/user"
)

func ToProjectDomain(req *CreateProjectRequest) *project.Project {
	return &project.Project{
		Name:      req.Name,
		Objective: req.Objective,
		CreatedBy: req.CreatedBy,
	}
}

func ToProjectResponse(proj *project.Project) *ProjectResponse {
	if proj == nil {
		return nil
	}

	response := &ProjectResponse{
		ID:        proj.ID,
		Name:      proj.Name,
		Objective: proj.Objective,
		CreatedBy: proj.CreatedBy,
		CreatedAt: proj.CreatedAt,
		UpdatedAt: proj.UpdatedAt,
	}

	if proj.Owner != nil {
		response.Owner = ToOwnerResponse(proj.Owner)
	}

	return response
}

func ToOwnerResponse(user *user.User) *OwnerResponse {
	if user == nil {
		return nil
	}

	return &OwnerResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToProjectListResponse(projects []project.Project) []ProjectResponse {
	responses := make([]ProjectResponse, len(projects))
	for i, proj := range projects {
		responses[i] = *ToProjectResponse(&proj)
	}
	return responses
}

package milestone

import (
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/project"
)

func ToMilestoneDomain(req *CreateMilestoneRequest) *milestone.Milestone {
	return &milestone.Milestone{
		ProjectID:   req.ProjectID,
		Title:       req.Title,
		Description: req.Description,
		ClassWeek:   req.ClassWeek,
	}
}

func ToMilestoneResponse(m *milestone.Milestone) *MilestoneResponse {
	if m == nil {
		return nil
	}

	response := &MilestoneResponse{
		ID:          m.ID,
		ProjectID:   m.ProjectID,
		Title:       m.Title,
		Description: m.Description,
		ClassWeek:   m.ClassWeek,
		CreatedAt:   m.CreatedAt,
	}

	if m.Project != nil {
		response.Project = ToProjectResponse(m.Project)
	}

	return response
}

func ToProjectResponse(p *project.Project) *ProjectResponse {
	if p == nil {
		return nil
	}

	return &ProjectResponse{
		ID:        p.ID,
		Name:      p.Name,
		Objective: p.Objective,
	}
}

func ToMilestoneListResponse(milestones []milestone.Milestone) []MilestoneResponse {
	responses := make([]MilestoneResponse, len(milestones))
	for i, m := range milestones {
		responses[i] = *ToMilestoneResponse(&m)
	}
	return responses
}

package deliverable

import (
	"softpharos/internal/core/domain/deliverable"
	"softpharos/internal/core/domain/milestone"
)

func ToDeliverableDomain(req *CreateDeliverableRequest) *deliverable.Deliverable {
	return &deliverable.Deliverable{
		MilestoneID: req.MilestoneID,
		URL:         req.URL,
		Type:        req.Type,
	}
}

func ToDeliverableResponse(d *deliverable.Deliverable) *DeliverableResponse {
	if d == nil {
		return nil
	}

	response := &DeliverableResponse{
		ID:          d.ID,
		MilestoneID: d.MilestoneID,
		URL:         d.URL,
		Type:        d.Type,
		CreatedAt:   d.CreatedAt,
	}

	if d.Milestone != nil {
		response.Milestone = ToMilestoneResponse(d.Milestone)
	}

	return response
}

func ToMilestoneResponse(m *milestone.Milestone) *MilestoneResponse {
	if m == nil {
		return nil
	}

	return &MilestoneResponse{
		ID:    m.ID,
		Title: m.Title,
	}
}

func ToDeliverableListResponse(deliverables []deliverable.Deliverable) []DeliverableResponse {
	responses := make([]DeliverableResponse, len(deliverables))
	for i, d := range deliverables {
		responses[i] = *ToDeliverableResponse(&d)
	}
	return responses
}

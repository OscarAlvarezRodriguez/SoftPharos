package reaction

import (
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/reaction"
	"softpharos/internal/core/domain/user"
)

func ToReactionDomain(req *CreateReactionRequest) *reaction.Reaction {
	return &reaction.Reaction{
		MilestoneID: req.MilestoneID,
		UserID:      req.UserID,
		Type:        req.Type,
	}
}

func ToReactionResponse(r *reaction.Reaction) *ReactionResponse {
	if r == nil {
		return nil
	}

	response := &ReactionResponse{
		ID:          r.ID,
		MilestoneID: r.MilestoneID,
		UserID:      r.UserID,
		Type:        r.Type,
		CreatedAt:   r.CreatedAt,
	}

	if r.Milestone != nil {
		response.Milestone = ToMilestoneResponse(r.Milestone)
	}

	if r.User != nil {
		response.User = ToUserResponse(r.User)
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

func ToReactionListResponse(reactions []reaction.Reaction) []ReactionResponse {
	responses := make([]ReactionResponse, len(reactions))
	for i, r := range reactions {
		responses[i] = *ToReactionResponse(&r)
	}
	return responses
}

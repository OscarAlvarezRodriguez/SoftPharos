package feedback

import (
	"softpharos/internal/core/domain/feedback"
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/user"
)

func ToFeedbackDomain(req *CreateFeedbackRequest) *feedback.Feedback {
	return &feedback.Feedback{
		MilestoneID: req.MilestoneID,
		ProfessorID: req.ProfessorID,
		Content:     req.Content,
	}
}

func ToFeedbackResponse(f *feedback.Feedback) *FeedbackResponse {
	if f == nil {
		return nil
	}

	response := &FeedbackResponse{
		ID:          f.ID,
		MilestoneID: f.MilestoneID,
		ProfessorID: f.ProfessorID,
		Content:     f.Content,
		CreatedAt:   f.CreatedAt,
	}

	if f.Milestone != nil {
		response.Milestone = ToMilestoneResponse(f.Milestone)
	}

	if f.Professor != nil {
		response.Professor = ToProfessorResponse(f.Professor)
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

func ToProfessorResponse(u *user.User) *ProfessorResponse {
	if u == nil {
		return nil
	}

	return &ProfessorResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func ToFeedbackListResponse(feedbacks []feedback.Feedback) []FeedbackResponse {
	responses := make([]FeedbackResponse, len(feedbacks))
	for i, f := range feedbacks {
		responses[i] = *ToFeedbackResponse(&f)
	}
	return responses
}

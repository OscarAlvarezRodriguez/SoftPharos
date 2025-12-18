package comment

import (
	"softpharos/internal/core/domain/comment"
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/user"
)

func ToCommentDomain(req *CreateCommentRequest) *comment.Comment {
	return &comment.Comment{
		MilestoneID: req.MilestoneID,
		UserID:      req.UserID,
		Content:     req.Content,
	}
}

func ToCommentResponse(c *comment.Comment) *CommentResponse {
	if c == nil {
		return nil
	}

	response := &CommentResponse{
		ID:          c.ID,
		MilestoneID: c.MilestoneID,
		UserID:      c.UserID,
		Content:     c.Content,
		CreatedAt:   c.CreatedAt,
	}

	if c.Milestone != nil {
		response.Milestone = ToMilestoneResponse(c.Milestone)
	}

	if c.User != nil {
		response.User = ToUserResponse(c.User)
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

func ToCommentListResponse(comments []comment.Comment) []CommentResponse {
	responses := make([]CommentResponse, len(comments))
	for i, c := range comments {
		responses[i] = *ToCommentResponse(&c)
	}
	return responses
}

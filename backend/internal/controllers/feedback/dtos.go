package feedback

import "time"

type CreateFeedbackRequest struct {
	MilestoneID int    `json:"milestone_id" binding:"required"`
	ProfessorID int    `json:"professor_id" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

type UpdateFeedbackRequest struct {
	Content string `json:"content" binding:"required"`
}

type FeedbackResponse struct {
	ID          int                `json:"id"`
	MilestoneID int                `json:"milestone_id"`
	Milestone   *MilestoneResponse `json:"milestone,omitempty"`
	ProfessorID int                `json:"professor_id"`
	Professor   *ProfessorResponse `json:"professor,omitempty"`
	Content     string             `json:"content"`
	CreatedAt   time.Time          `json:"created_at"`
}

type MilestoneResponse struct {
	ID    int     `json:"id"`
	Title *string `json:"title"`
}

type ProfessorResponse struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

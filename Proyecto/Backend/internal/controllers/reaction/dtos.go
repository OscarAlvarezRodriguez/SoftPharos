package reaction

import "time"

type CreateReactionRequest struct {
	MilestoneID int     `json:"milestone_id" binding:"required"`
	UserID      int     `json:"user_id" binding:"required"`
	Type        *string `json:"type"`
}

type UpdateReactionRequest struct {
	Type *string `json:"type"`
}

type ReactionResponse struct {
	ID          int                `json:"id"`
	MilestoneID int                `json:"milestone_id"`
	Milestone   *MilestoneResponse `json:"milestone,omitempty"`
	UserID      int                `json:"user_id"`
	User        *UserResponse      `json:"user,omitempty"`
	Type        *string            `json:"type"`
	CreatedAt   time.Time          `json:"created_at"`
}

type MilestoneResponse struct {
	ID    int     `json:"id"`
	Title *string `json:"title"`
}

type UserResponse struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

package comment

import "time"

type CreateCommentRequest struct {
	MilestoneID int     `json:"milestone_id" binding:"required"`
	UserID      int     `json:"user_id" binding:"required"`
	Content     *string `json:"content"`
}

type UpdateCommentRequest struct {
	Content *string `json:"content"`
}

type CommentResponse struct {
	ID          int                `json:"id"`
	MilestoneID int                `json:"milestone_id"`
	Milestone   *MilestoneResponse `json:"milestone,omitempty"`
	UserID      int                `json:"user_id"`
	User        *UserResponse      `json:"user,omitempty"`
	Content     *string            `json:"content"`
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

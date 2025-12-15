package deliverable

import "time"

type CreateDeliverableRequest struct {
	MilestoneID int     `json:"milestone_id" binding:"required"`
	URL         string  `json:"url" binding:"required"`
	Type        *string `json:"type"`
}

type UpdateDeliverableRequest struct {
	URL  *string `json:"url"`
	Type *string `json:"type"`
}

type DeliverableResponse struct {
	ID          int                `json:"id"`
	MilestoneID int                `json:"milestone_id"`
	Milestone   *MilestoneResponse `json:"milestone,omitempty"`
	URL         string             `json:"url"`
	Type        *string            `json:"type"`
	CreatedAt   time.Time          `json:"created_at"`
}

type MilestoneResponse struct {
	ID    int     `json:"id"`
	Title *string `json:"title"`
}

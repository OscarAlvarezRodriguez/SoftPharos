package milestone

import "time"

type CreateMilestoneRequest struct {
	ProjectID   int     `json:"project_id" binding:"required"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	ClassWeek   *int    `json:"class_week"`
}

type UpdateMilestoneRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	ClassWeek   *int    `json:"class_week"`
}

type MilestoneResponse struct {
	ID          int              `json:"id"`
	ProjectID   int              `json:"project_id"`
	Project     *ProjectResponse `json:"project,omitempty"`
	Title       *string          `json:"title"`
	Description *string          `json:"description"`
	ClassWeek   *int             `json:"class_week"`
	CreatedAt   time.Time        `json:"created_at"`
}

type ProjectResponse struct {
	ID        int     `json:"id"`
	Name      *string `json:"name"`
	Objective *string `json:"objective"`
}

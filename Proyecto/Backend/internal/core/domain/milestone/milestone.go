package milestone

import (
	"time"

	"softpharos/internal/core/domain/project"
)

type Milestone struct {
	ID          int
	ProjectID   int
	Project     *project.Project
	Title       *string
	Description *string
	ClassWeek   *int
	CreatedAt   time.Time
}

package deliverable

import (
	"time"

	"softpharos/internal/core/domain/milestone"
)

type Deliverable struct {
	ID          int
	MilestoneID int
	Milestone   *milestone.Milestone
	URL         string
	Type        *string
	CreatedAt   time.Time
}

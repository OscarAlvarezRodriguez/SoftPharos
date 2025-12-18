package reaction

import (
	"time"

	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/user"
)

type Reaction struct {
	ID          int
	MilestoneID int
	Milestone   *milestone.Milestone
	UserID      int
	User        *user.User
	Type        *string
	CreatedAt   time.Time
}

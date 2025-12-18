package comment

import (
	"time"

	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/user"
)

type Comment struct {
	ID          int
	MilestoneID int
	Milestone   *milestone.Milestone
	UserID      int
	User        *user.User
	Content     *string
	CreatedAt   time.Time
}

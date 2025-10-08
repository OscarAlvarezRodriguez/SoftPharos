package feedback

import (
	"time"

	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/user"
)

type Feedback struct {
	ID          int
	MilestoneID int
	Milestone   *milestone.Milestone
	ProfessorID int
	Professor   *user.User
	Content     string
	CreatedAt   time.Time
}

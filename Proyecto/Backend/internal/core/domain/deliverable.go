package domain

import "time"

// Deliverable representa un entregable de un hito
type Deliverable struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	MilestoneID int        `json:"milestone_id" gorm:"not null"`
	Milestone   *Milestone `json:"milestone,omitempty" gorm:"foreignKey:MilestoneID"`
	URL         string     `json:"url" gorm:"type:text;not null"`
	Type        *string    `json:"type" gorm:"type:varchar"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (Deliverable) TableName() string {
	return "deliverables"
}

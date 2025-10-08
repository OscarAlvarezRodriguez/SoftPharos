package domain

import "time"

// Reaction representa una reacción de un usuario a un hito
type Reaction struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	MilestoneID int        `json:"milestone_id" gorm:"not null"`
	Milestone   *Milestone `json:"milestone,omitempty" gorm:"foreignKey:MilestoneID"`
	UserID      int        `json:"user_id" gorm:"not null"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Type        *string    `json:"type" gorm:"type:varchar"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (Reaction) TableName() string {
	return "reactions"
}

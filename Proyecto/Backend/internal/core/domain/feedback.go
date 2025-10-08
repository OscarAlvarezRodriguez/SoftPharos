package domain

import "time"

// Feedback representa la retroalimentación de un profesor sobre un hito
type Feedback struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	MilestoneID int        `json:"milestone_id" gorm:"not null"`
	Milestone   *Milestone `json:"milestone,omitempty" gorm:"foreignKey:MilestoneID"`
	ProfessorID int        `json:"professor_id" gorm:"not null"`
	Professor   *User      `json:"professor,omitempty" gorm:"foreignKey:ProfessorID"`
	Content     string     `json:"content" gorm:"type:text;not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (Feedback) TableName() string {
	return "feedback"
}

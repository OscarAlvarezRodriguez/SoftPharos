package domain

import "time"

// Milestone representa un hito de un proyecto
type Milestone struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID   int       `json:"project_id" gorm:"not null"`
	Project     *Project  `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	Title       *string   `json:"title" gorm:"type:varchar"`
	Description *string   `json:"description" gorm:"type:text"`
	ClassWeek   *int      `json:"class_week" gorm:"type:integer"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (Milestone) TableName() string {
	return "milestones"
}

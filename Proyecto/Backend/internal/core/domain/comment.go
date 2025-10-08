package domain

import "time"

// Comment representa un comentario en un hito
type Comment struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	MilestoneID int        `json:"milestone_id" gorm:"not null"`
	Milestone   *Milestone `json:"milestone,omitempty" gorm:"foreignKey:MilestoneID"`
	UserID      int        `json:"user_id" gorm:"not null"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Content     *string    `json:"content" gorm:"type:text"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (Comment) TableName() string {
	return "comments"
}

package models

import "time"

type FeedbackModel struct {
	ID          int             `gorm:"primaryKey;autoIncrement"`
	MilestoneID int             `gorm:"not null"`
	Milestone   *MilestoneModel `gorm:"foreignKey:MilestoneID"`
	ProfessorID int             `gorm:"not null"`
	Professor   *UserModel      `gorm:"foreignKey:ProfessorID"`
	Content     string          `gorm:"type:text;not null"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
}

func (FeedbackModel) TableName() string {
	return "feedback"
}

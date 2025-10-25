package models

import "time"

type MilestoneModel struct {
	ID          int           `gorm:"primaryKey;autoIncrement"`
	ProjectID   int           `gorm:"not null"`
	Project     *ProjectModel `gorm:"foreignKey:ProjectID"`
	Title       *string       `gorm:"type:varchar"`
	Description *string       `gorm:"type:text"`
	ClassWeek   *int          `gorm:"type:integer"`
	CreatedAt   time.Time     `gorm:"autoCreateTime"`
}

func (MilestoneModel) TableName() string {
	return "milestone"
}

package models

import "time"

type ProjectModel struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	Name      *string    `gorm:"type:varchar"`
	Objective *string    `gorm:"type:text"`
	CreatedBy int        `gorm:"not null"`
	Creator   *UserModel `gorm:"foreignKey:CreatedBy"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

func (ProjectModel) TableName() string {
	return "project"
}

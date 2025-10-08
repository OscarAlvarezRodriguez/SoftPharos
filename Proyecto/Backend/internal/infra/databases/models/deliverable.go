package models

import "time"

type DeliverableModel struct {
	ID          int             `gorm:"primaryKey;autoIncrement"`
	MilestoneID int             `gorm:"not null"`
	Milestone   *MilestoneModel `gorm:"foreignKey:MilestoneID"`
	URL         string          `gorm:"type:text;not null"`
	Type        *string         `gorm:"type:varchar"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
}

func (DeliverableModel) TableName() string {
	return "deliverables"
}

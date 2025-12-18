package models

import "time"

type ReactionModel struct {
	ID          int             `gorm:"primaryKey;autoIncrement"`
	MilestoneID int             `gorm:"not null"`
	Milestone   *MilestoneModel `gorm:"foreignKey:MilestoneID"`
	UserID      int             `gorm:"not null"`
	User        *UserModel      `gorm:"foreignKey:UserID"`
	Type        *string         `gorm:"type:varchar"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
}

func (ReactionModel) TableName() string {
	return "reaction"
}

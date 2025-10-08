package models

import "time"

type CommentModel struct {
	ID          int             `gorm:"primaryKey;autoIncrement"`
	MilestoneID int             `gorm:"not null"`
	Milestone   *MilestoneModel `gorm:"foreignKey:MilestoneID"`
	UserID      int             `gorm:"not null"`
	User        *UserModel      `gorm:"foreignKey:UserID"`
	Content     *string         `gorm:"type:text"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
}

func (CommentModel) TableName() string {
	return "comments"
}

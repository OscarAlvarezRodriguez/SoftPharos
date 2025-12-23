package models

import "time"

type UserModel struct {
	ID         int        `gorm:"primaryKey;autoIncrement"`
	Name       *string    `gorm:"type:varchar"`
	Email      string     `gorm:"unique;not null"`
	ProviderID string     `gorm:"type:varchar;unique;not null"`
	RoleID     int        `gorm:"not null"`
	Role       *RoleModel `gorm:"foreignKey:RoleID"`
	PictureURL *string    `gorm:"type:varchar"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
}

func (UserModel) TableName() string {
	return "user"
}

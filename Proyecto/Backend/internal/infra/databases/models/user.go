package models

import "time"

type UserModel struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	Name      *string    `gorm:"type:varchar"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	RoleID    int        `gorm:"not null"`
	Role      *RoleModel `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
}

func (UserModel) TableName() string {
	return "user"
}

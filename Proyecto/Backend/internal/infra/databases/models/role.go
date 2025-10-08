package models

import "time"

type RoleModel struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"unique;not null"`
	Description *string   `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (RoleModel) TableName() string {
	return "roles"
}

package domain

import "time"

// Role representa el rol de un usuario en el sistema
type Role struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"unique;not null"`
	Description *string   `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (Role) TableName() string {
	return "roles"
}

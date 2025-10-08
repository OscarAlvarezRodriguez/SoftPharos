package domain

import "time"

// User representa un usuario del sistema
type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      *string   `json:"name" gorm:"type:varchar"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // No exponer en JSON
	RoleID    int       `json:"role_id" gorm:"not null"`
	Role      *Role     `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (User) TableName() string {
	return "users"
}

package domain

import "time"

// Project representa un proyecto en el sistema
type Project struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      *string   `json:"name" gorm:"type:varchar"`
	Objective *string   `json:"objective" gorm:"type:text"`
	CreatedBy int       `json:"created_by" gorm:"not null"`
	Creator   *User     `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (Project) TableName() string {
	return "projects"
}

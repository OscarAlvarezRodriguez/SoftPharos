package domain

import "time"

// ProjectMember representa un miembro de un proyecto
type ProjectMember struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID int       `json:"project_id" gorm:"not null"`
	Project   *Project  `json:"project,omitempty" gorm:"foreignKey:ProjectID"`
	UserID    int       `json:"user_id" gorm:"not null"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Role      *string   `json:"role" gorm:"type:varchar"`
	JoinedAt  time.Time `json:"joined_at" gorm:"autoCreateTime"`
}

// TableName especifica el nombre de la tabla en la BD
func (ProjectMember) TableName() string {
	return "project_members"
}

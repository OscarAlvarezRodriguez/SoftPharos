package models

import "time"

type ProjectMemberModel struct {
	ID        int           `gorm:"primaryKey;autoIncrement"`
	ProjectID int           `gorm:"not null"`
	Project   *ProjectModel `gorm:"foreignKey:ProjectID"`
	UserID    int           `gorm:"not null"`
	User      *UserModel    `gorm:"foreignKey:UserID"`
	Role      *string       `gorm:"type:varchar"`
	JoinedAt  time.Time     `gorm:"autoCreateTime"`
}

func (ProjectMemberModel) TableName() string {
	return "project_member"
}

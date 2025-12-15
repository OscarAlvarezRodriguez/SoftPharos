package mappers

import (
	"softpharos/internal/core/domain/project_member"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProjectMemberToDomain(t *testing.T) {
	role := "Developer"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.ProjectMemberModel
		expected *project_member.ProjectMember
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.ProjectMemberModel{
				ID:        1,
				ProjectID: 1,
				UserID:    1,
				Role:      &role,
				JoinedAt:  now,
			},
			expected: &project_member.ProjectMember{
				ID:        1,
				ProjectID: 1,
				UserID:    1,
				Role:      &role,
				JoinedAt:  now,
			},
		},
		{
			name:     "retorna nil para modelo nil",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProjectMemberToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProjectMemberToModel(t *testing.T) {
	role := "Developer"
	now := time.Now()

	tests := []struct {
		name     string
		input    *project_member.ProjectMember
		expected *models.ProjectMemberModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &project_member.ProjectMember{
				ID:        1,
				ProjectID: 1,
				UserID:    1,
				Role:      &role,
				JoinedAt:  now,
			},
			expected: &models.ProjectMemberModel{
				ID:        1,
				ProjectID: 1,
				UserID:    1,
				Role:      &role,
				JoinedAt:  now,
			},
		},
		{
			name:     "retorna nil para dominio nil",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProjectMemberToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProjectMemberListToDomain(t *testing.T) {
	role := "Developer"

	tests := []struct {
		name     string
		input    []models.ProjectMemberModel
		expected []project_member.ProjectMember
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.ProjectMemberModel{
				{ID: 1, ProjectID: 1, UserID: 1, Role: &role},
				{ID: 2, ProjectID: 1, UserID: 2, Role: &role},
			},
			expected: []project_member.ProjectMember{
				{ID: 1, ProjectID: 1, UserID: 1, Role: &role},
				{ID: 2, ProjectID: 1, UserID: 2, Role: &role},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.ProjectMemberModel{},
			expected: []project_member.ProjectMember{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProjectMemberListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

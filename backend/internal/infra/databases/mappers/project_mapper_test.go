package mappers

import (
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/domain/user"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProjectToDomain(t *testing.T) {
	name := "Test Project"
	objective := "Test Objective"
	userName := "John Doe"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.ProjectModel
		expected *project.Project
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.ProjectModel{
				ID:        1,
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
				Owner: &models.UserModel{
					ID:    1,
					Name:  &userName,
					Email: "john@example.com",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expected: &project.Project{
				ID:        1,
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
				Owner: &user.User{
					ID:    1,
					Name:  &userName,
					Email: "john@example.com",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			name:     "retorna nil para modelo nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja owner nil",
			input: &models.ProjectModel{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
				Owner:     nil,
			},
			expected: &project.Project{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
				Owner:     nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProjectToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProjectToModel(t *testing.T) {
	name := "Test Project"
	objective := "Test Objective"
	userName := "John Doe"
	now := time.Now()

	tests := []struct {
		name     string
		input    *project.Project
		expected *models.ProjectModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &project.Project{
				ID:        1,
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
				Owner: &user.User{
					ID:    1,
					Name:  &userName,
					Email: "john@example.com",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expected: &models.ProjectModel{
				ID:        1,
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
				Owner: &models.UserModel{
					ID:    1,
					Name:  &userName,
					Email: "john@example.com",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			name:     "retorna nil para dominio nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja owner nil",
			input: &project.Project{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
				Owner:     nil,
			},
			expected: &models.ProjectModel{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
				Owner:     nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProjectToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProjectListToDomain(t *testing.T) {
	name1 := "Project 1"
	name2 := "Project 2"

	tests := []struct {
		name     string
		input    []models.ProjectModel
		expected []project.Project
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.ProjectModel{
				{ID: 1, Name: &name1, CreatedBy: 1},
				{ID: 2, Name: &name2, CreatedBy: 2},
			},
			expected: []project.Project{
				{ID: 1, Name: &name1, CreatedBy: 1},
				{ID: 2, Name: &name2, CreatedBy: 2},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.ProjectModel{},
			expected: []project.Project{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProjectListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

package mappers

import (
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/project"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMilestoneToDomain(t *testing.T) {
	title := "Sprint 1"
	description := "First sprint"
	classWeek := 1
	projectName := "Test Project"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.MilestoneModel
		expected *milestone.Milestone
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.MilestoneModel{
				ID:          1,
				ProjectID:   1,
				Title:       &title,
				Description: &description,
				ClassWeek:   &classWeek,
				Project: &models.ProjectModel{
					ID:   1,
					Name: &projectName,
				},
				CreatedAt: now,
			},
			expected: &milestone.Milestone{
				ID:          1,
				ProjectID:   1,
				Title:       &title,
				Description: &description,
				ClassWeek:   &classWeek,
				Project: &project.Project{
					ID:   1,
					Name: &projectName,
				},
				CreatedAt: now,
			},
		},
		{
			name:     "retorna nil para modelo nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja project nil",
			input: &models.MilestoneModel{
				ID:        1,
				ProjectID: 1,
				Title:     &title,
				Project:   nil,
			},
			expected: &milestone.Milestone{
				ID:        1,
				ProjectID: 1,
				Title:     &title,
				Project:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MilestoneToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMilestoneToModel(t *testing.T) {
	title := "Sprint 1"
	description := "First sprint"
	classWeek := 1
	projectName := "Test Project"
	now := time.Now()

	tests := []struct {
		name     string
		input    *milestone.Milestone
		expected *models.MilestoneModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &milestone.Milestone{
				ID:          1,
				ProjectID:   1,
				Title:       &title,
				Description: &description,
				ClassWeek:   &classWeek,
				Project: &project.Project{
					ID:   1,
					Name: &projectName,
				},
				CreatedAt: now,
			},
			expected: &models.MilestoneModel{
				ID:          1,
				ProjectID:   1,
				Title:       &title,
				Description: &description,
				ClassWeek:   &classWeek,
				Project: &models.ProjectModel{
					ID:   1,
					Name: &projectName,
				},
				CreatedAt: now,
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
			result := MilestoneToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMilestoneListToDomain(t *testing.T) {
	title1 := "Milestone 1"
	title2 := "Milestone 2"

	tests := []struct {
		name     string
		input    []models.MilestoneModel
		expected []milestone.Milestone
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.MilestoneModel{
				{ID: 1, ProjectID: 1, Title: &title1},
				{ID: 2, ProjectID: 1, Title: &title2},
			},
			expected: []milestone.Milestone{
				{ID: 1, ProjectID: 1, Title: &title1},
				{ID: 2, ProjectID: 1, Title: &title2},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.MilestoneModel{},
			expected: []milestone.Milestone{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MilestoneListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

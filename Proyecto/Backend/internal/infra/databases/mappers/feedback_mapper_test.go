package mappers

import (
	"softpharos/internal/core/domain/feedback"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFeedbackToDomain(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.FeedbackModel
		expected *feedback.Feedback
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.FeedbackModel{
				ID:          1,
				MilestoneID: 1,
				ProfessorID: 1,
				Content:     "Good work",
				CreatedAt:   now,
			},
			expected: &feedback.Feedback{
				ID:          1,
				MilestoneID: 1,
				ProfessorID: 1,
				Content:     "Good work",
				CreatedAt:   now,
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
			result := FeedbackToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFeedbackToModel(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		input    *feedback.Feedback
		expected *models.FeedbackModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &feedback.Feedback{
				ID:          1,
				MilestoneID: 1,
				ProfessorID: 1,
				Content:     "Good work",
				CreatedAt:   now,
			},
			expected: &models.FeedbackModel{
				ID:          1,
				MilestoneID: 1,
				ProfessorID: 1,
				Content:     "Good work",
				CreatedAt:   now,
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
			result := FeedbackToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFeedbackListToDomain(t *testing.T) {
	tests := []struct {
		name     string
		input    []models.FeedbackModel
		expected []feedback.Feedback
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.FeedbackModel{
				{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Good"},
				{ID: 2, MilestoneID: 1, ProfessorID: 1, Content: "Excellent"},
			},
			expected: []feedback.Feedback{
				{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Good"},
				{ID: 2, MilestoneID: 1, ProfessorID: 1, Content: "Excellent"},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.FeedbackModel{},
			expected: []feedback.Feedback{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FeedbackListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

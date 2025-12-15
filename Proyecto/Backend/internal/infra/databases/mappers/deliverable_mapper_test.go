package mappers

import (
	"softpharos/internal/core/domain/deliverable"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeliverableToDomain(t *testing.T) {
	typeVal := "PDF"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.DeliverableModel
		expected *deliverable.Deliverable
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.DeliverableModel{
				ID:          1,
				MilestoneID: 1,
				URL:         "http://example.com",
				Type:        &typeVal,
				CreatedAt:   now,
			},
			expected: &deliverable.Deliverable{
				ID:          1,
				MilestoneID: 1,
				URL:         "http://example.com",
				Type:        &typeVal,
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
			result := DeliverableToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDeliverableToModel(t *testing.T) {
	typeVal := "PDF"
	now := time.Now()

	tests := []struct {
		name     string
		input    *deliverable.Deliverable
		expected *models.DeliverableModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &deliverable.Deliverable{
				ID:          1,
				MilestoneID: 1,
				URL:         "http://example.com",
				Type:        &typeVal,
				CreatedAt:   now,
			},
			expected: &models.DeliverableModel{
				ID:          1,
				MilestoneID: 1,
				URL:         "http://example.com",
				Type:        &typeVal,
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
			result := DeliverableToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDeliverableListToDomain(t *testing.T) {
	tests := []struct {
		name     string
		input    []models.DeliverableModel
		expected []deliverable.Deliverable
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.DeliverableModel{
				{ID: 1, MilestoneID: 1, URL: "http://example.com/1"},
				{ID: 2, MilestoneID: 1, URL: "http://example.com/2"},
			},
			expected: []deliverable.Deliverable{
				{ID: 1, MilestoneID: 1, URL: "http://example.com/1"},
				{ID: 2, MilestoneID: 1, URL: "http://example.com/2"},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.DeliverableModel{},
			expected: []deliverable.Deliverable{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeliverableListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

package mappers

import (
	"softpharos/internal/core/domain/reaction"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReactionToDomain(t *testing.T) {
	reactionType := "like"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.ReactionModel
		expected *reaction.Reaction
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.ReactionModel{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Type:        &reactionType,
				CreatedAt:   now,
			},
			expected: &reaction.Reaction{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Type:        &reactionType,
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
			result := ReactionToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReactionToModel(t *testing.T) {
	reactionType := "like"
	now := time.Now()

	tests := []struct {
		name     string
		input    *reaction.Reaction
		expected *models.ReactionModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &reaction.Reaction{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Type:        &reactionType,
				CreatedAt:   now,
			},
			expected: &models.ReactionModel{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Type:        &reactionType,
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
			result := ReactionToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReactionListToDomain(t *testing.T) {
	reactionType := "like"

	tests := []struct {
		name     string
		input    []models.ReactionModel
		expected []reaction.Reaction
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.ReactionModel{
				{ID: 1, MilestoneID: 1, UserID: 1, Type: &reactionType},
				{ID: 2, MilestoneID: 1, UserID: 2, Type: &reactionType},
			},
			expected: []reaction.Reaction{
				{ID: 1, MilestoneID: 1, UserID: 1, Type: &reactionType},
				{ID: 2, MilestoneID: 1, UserID: 2, Type: &reactionType},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.ReactionModel{},
			expected: []reaction.Reaction{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReactionListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

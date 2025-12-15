package mappers

import (
	"softpharos/internal/core/domain/comment"
	"softpharos/internal/core/domain/milestone"
	"softpharos/internal/core/domain/user"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCommentToDomain(t *testing.T) {
	content := "Great work"
	milestoneTitle := "Sprint 1"
	userName := "John Doe"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.CommentModel
		expected *comment.Comment
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.CommentModel{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Content:     &content,
				Milestone:   &models.MilestoneModel{ID: 1, Title: &milestoneTitle},
				User:        &models.UserModel{ID: 1, Name: &userName, Email: "john@example.com"},
				CreatedAt:   now,
			},
			expected: &comment.Comment{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Content:     &content,
				Milestone:   &milestone.Milestone{ID: 1, Title: &milestoneTitle},
				User:        &user.User{ID: 1, Name: &userName, Email: "john@example.com"},
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
			result := CommentToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCommentToModel(t *testing.T) {
	content := "Great work"
	now := time.Now()

	tests := []struct {
		name     string
		input    *comment.Comment
		expected *models.CommentModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &comment.Comment{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Content:     &content,
				CreatedAt:   now,
			},
			expected: &models.CommentModel{
				ID:          1,
				MilestoneID: 1,
				UserID:      1,
				Content:     &content,
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
			result := CommentToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCommentListToDomain(t *testing.T) {
	content1 := "Comment 1"
	content2 := "Comment 2"

	tests := []struct {
		name     string
		input    []models.CommentModel
		expected []comment.Comment
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.CommentModel{
				{ID: 1, MilestoneID: 1, UserID: 1, Content: &content1},
				{ID: 2, MilestoneID: 1, UserID: 1, Content: &content2},
			},
			expected: []comment.Comment{
				{ID: 1, MilestoneID: 1, UserID: 1, Content: &content1},
				{ID: 2, MilestoneID: 1, UserID: 1, Content: &content2},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.CommentModel{},
			expected: []comment.Comment{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CommentListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

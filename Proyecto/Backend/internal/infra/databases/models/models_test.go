package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllModelsTableNames(t *testing.T) {
	tests := []struct {
		name          string
		model         interface{ TableName() string }
		expectedTable string
	}{
		{"User", UserModel{}, "user"},
		{"Role", RoleModel{}, "role"},
		{"Project", ProjectModel{}, "project"},
		{"Milestone", MilestoneModel{}, "milestone"},
		{"Comment", CommentModel{}, "comment"},
		{"Deliverable", DeliverableModel{}, "deliverable"},
		{"Feedback", FeedbackModel{}, "feedback"},
		{"ProjectMember", ProjectMemberModel{}, "project_member"},
		{"Reaction", ReactionModel{}, "reaction"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tableName := tt.model.TableName()
			assert.Equal(t, tt.expectedTable, tableName,
				"El nombre de tabla para %s debe ser '%s'", tt.name, tt.expectedTable)
		})
	}
}

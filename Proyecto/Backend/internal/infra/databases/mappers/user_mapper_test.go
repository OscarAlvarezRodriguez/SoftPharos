package mappers

import (
	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/domain/user"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserToDomain(t *testing.T) {
	name := "John Doe"
	roleName := "Admin"
	roleDesc := "Administrator"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.UserModel
		expected *user.User
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.UserModel{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role: &models.RoleModel{
					ID:          1,
					Name:        roleName,
					Description: &roleDesc,
				},
				CreatedAt: now,
			},
			expected: &user.User{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role: &role.Role{
					ID:          1,
					Name:        roleName,
					Description: &roleDesc,
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
			name: "maneja role nil",
			input: &models.UserModel{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role:     nil,
			},
			expected: &user.User{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role:     nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UserToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUserToModel(t *testing.T) {
	name := "John Doe"
	roleName := "Admin"
	roleDesc := "Administrator"
	now := time.Now()

	tests := []struct {
		name     string
		input    *user.User
		expected *models.UserModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &user.User{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role: &role.Role{
					ID:          1,
					Name:        roleName,
					Description: &roleDesc,
				},
				CreatedAt: now,
			},
			expected: &models.UserModel{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role: &models.RoleModel{
					ID:          1,
					Name:        roleName,
					Description: &roleDesc,
				},
				CreatedAt: now,
			},
		},
		{
			name:     "retorna nil para dominio nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja role nil",
			input: &user.User{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role:     nil,
			},
			expected: &models.UserModel{
				ID:       1,
				Name:     &name,
				Email:    "john@example.com",
				Password: "hash",
				RoleID:   1,
				Role:     nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UserToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUserListToDomain(t *testing.T) {
	name1 := "User 1"
	name2 := "User 2"

	tests := []struct {
		name     string
		input    []models.UserModel
		expected []user.User
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.UserModel{
				{ID: 1, Name: &name1, Email: "user1@example.com", Password: "hash1", RoleID: 1},
				{ID: 2, Name: &name2, Email: "user2@example.com", Password: "hash2", RoleID: 2},
			},
			expected: []user.User{
				{ID: 1, Name: &name1, Email: "user1@example.com", Password: "hash1", RoleID: 1},
				{ID: 2, Name: &name2, Email: "user2@example.com", Password: "hash2", RoleID: 2},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.UserModel{},
			expected: []user.User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UserListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

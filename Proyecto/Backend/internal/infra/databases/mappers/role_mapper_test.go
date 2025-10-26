package mappers

import (
	"softpharos/internal/core/domain/role"
	"softpharos/internal/infra/databases/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoleToDomain(t *testing.T) {
	description := "Admin role description"
	now := time.Now()

	tests := []struct {
		name     string
		input    *models.RoleModel
		expected *role.Role
	}{
		{
			name: "convierte modelo válido a dominio",
			input: &models.RoleModel{
				ID:          1,
				Name:        "Admin",
				Description: &description,
				CreatedAt:   now,
			},
			expected: &role.Role{
				ID:          1,
				Name:        "Admin",
				Description: &description,
				CreatedAt:   now,
			},
		},
		{
			name:     "retorna nil para modelo nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja description nil",
			input: &models.RoleModel{
				ID:          1,
				Name:        "Admin",
				Description: nil,
			},
			expected: &role.Role{
				ID:          1,
				Name:        "Admin",
				Description: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RoleToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRoleToModel(t *testing.T) {
	description := "Admin role description"
	now := time.Now()

	tests := []struct {
		name     string
		input    *role.Role
		expected *models.RoleModel
	}{
		{
			name: "convierte dominio válido a modelo",
			input: &role.Role{
				ID:          1,
				Name:        "Admin",
				Description: &description,
				CreatedAt:   now,
			},
			expected: &models.RoleModel{
				ID:          1,
				Name:        "Admin",
				Description: &description,
				CreatedAt:   now,
			},
		},
		{
			name:     "retorna nil para dominio nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja description nil",
			input: &role.Role{
				ID:          1,
				Name:        "Admin",
				Description: nil,
			},
			expected: &models.RoleModel{
				ID:          1,
				Name:        "Admin",
				Description: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RoleToModel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRoleListToDomain(t *testing.T) {
	desc1 := "Admin description"
	desc2 := "User description"

	tests := []struct {
		name     string
		input    []models.RoleModel
		expected []role.Role
	}{
		{
			name: "convierte lista de modelos a dominios",
			input: []models.RoleModel{
				{ID: 1, Name: "Admin", Description: &desc1},
				{ID: 2, Name: "User", Description: &desc2},
			},
			expected: []role.Role{
				{ID: 1, Name: "Admin", Description: &desc1},
				{ID: 2, Name: "User", Description: &desc2},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []models.RoleModel{},
			expected: []role.Role{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RoleListToDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

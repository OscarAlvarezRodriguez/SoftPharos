package role

import (
	"softpharos/internal/core/domain/role"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToRoleResponse(t *testing.T) {
	desc := "Administrator"
	now := time.Now()

	tests := []struct {
		name     string
		input    *role.Role
		expected *RoleResponse
	}{
		{
			name: "convierte dominio válido a response",
			input: &role.Role{
				ID:          1,
				Name:        "Admin",
				Description: &desc,
				CreatedAt:   now,
			},
			expected: &RoleResponse{
				ID:          1,
				Name:        "Admin",
				Description: &desc,
				CreatedAt:   now,
			},
		},
		{
			name:     "retorna nil para rol nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja description nil",
			input: &role.Role{
				ID:          1,
				Name:        "Admin",
				Description: nil,
				CreatedAt:   now,
			},
			expected: &RoleResponse{
				ID:          1,
				Name:        "Admin",
				Description: nil,
				CreatedAt:   now,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToRoleResponse(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToRoleListResponse(t *testing.T) {
	desc1 := "Administrator"
	desc2 := "Developer"
	now := time.Now()

	tests := []struct {
		name     string
		input    []role.Role
		expected []RoleResponse
	}{
		{
			name: "convierte lista de dominios a responses",
			input: []role.Role{
				{ID: 1, Name: "Admin", Description: &desc1, CreatedAt: now},
				{ID: 2, Name: "Developer", Description: &desc2, CreatedAt: now},
			},
			expected: []RoleResponse{
				{ID: 1, Name: "Admin", Description: &desc1, CreatedAt: now},
				{ID: 2, Name: "Developer", Description: &desc2, CreatedAt: now},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []role.Role{},
			expected: []RoleResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToRoleListResponse(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

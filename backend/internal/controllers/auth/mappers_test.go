package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/domain/user"
)

func TestToUserInfo(t *testing.T) {
	name := "John Doe"
	picture := "https://example.com/pic.jpg"
	roleName := "student"
	roleDesc := "Student role"
	now := time.Now()

	tests := []struct {
		name        string
		user        *user.User
		accessToken string
		expected    *AuthResponse
	}{
		{
			name: "convierte usuario con rol a UserInfo",
			user: &user.User{
				ID:         1,
				Name:       &name,
				Email:      "test@unal.edu.co",
				ProviderID: "google-123",
				RoleID:     3,
				PictureURL: &picture,
				Role: &role.Role{
					ID:          3,
					Name:        roleName,
					Description: &roleDesc,
				},
				CreatedAt: now,
			},
			accessToken: "jwt-token-123",
			expected: &AuthResponse{
				AccessToken: "jwt-token-123",
				User: UserInfo{
					ID:         1,
					Name:       &name,
					Email:      "test@unal.edu.co",
					RoleID:     3,
					RoleName:   roleName,
					PictureURL: &picture,
					CreatedAt:  now,
				},
			},
		},
		{
			name: "convierte usuario sin rol (usa default)",
			user: &user.User{
				ID:         2,
				Name:       &name,
				Email:      "test2@unal.edu.co",
				ProviderID: "google-456",
				RoleID:     3,
				Role:       nil,
				CreatedAt:  now,
			},
			accessToken: "jwt-token-456",
			expected: &AuthResponse{
				AccessToken: "jwt-token-456",
				User: UserInfo{
					ID:        2,
					Name:      &name,
					Email:     "test2@unal.edu.co",
					RoleID:    3,
					RoleName:  "student",
					CreatedAt: now,
				},
			},
		},
		{
			name:        "retorna nil para usuario nil",
			user:        nil,
			accessToken: "jwt-token",
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToUserInfo(tt.user, tt.accessToken)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToRoleResponse(t *testing.T) {
	desc := "Test role"

	tests := []struct {
		name     string
		role     *role.Role
		expected *RoleInfo
	}{
		{
			name: "convierte rol válido",
			role: &role.Role{
				ID:          1,
				Name:        "admin",
				Description: &desc,
			},
			expected: &RoleInfo{
				ID:          1,
				Name:        "admin",
				Description: &desc,
			},
		},
		{
			name:     "retorna nil para rol nil",
			role:     nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToRoleResponse(tt.role)
			assert.Equal(t, tt.expected, result)
		})
	}
}

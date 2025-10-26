package project

import (
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/domain/user"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToProjectDomain(t *testing.T) {
	name := "Test Project"
	objective := "Test Objective"

	tests := []struct {
		name     string
		input    *CreateProjectRequest
		expected *project.Project
	}{
		{
			name: "convierte request válida a dominio",
			input: &CreateProjectRequest{
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
			},
			expected: &project.Project{
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
			},
		},
		{
			name: "maneja objective nil",
			input: &CreateProjectRequest{
				Name:      &name,
				Objective: nil,
				CreatedBy: 1,
			},
			expected: &project.Project{
				Name:      &name,
				Objective: nil,
				CreatedBy: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToProjectDomain(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToProjectResponse(t *testing.T) {
	name := "Test Project"
	objective := "Test Objective"
	userName := "John Doe"
	now := time.Now()

	tests := []struct {
		name     string
		input    *project.Project
		expected *ProjectResponse
	}{
		{
			name: "convierte dominio válido a response",
			input: &project.Project{
				ID:        1,
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
				Owner: &user.User{
					ID:    1,
					Name:  &userName,
					Email: "john@example.com",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expected: &ProjectResponse{
				ID:        1,
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
				Owner: &OwnerResponse{
					ID:    1,
					Name:  &userName,
					Email: "john@example.com",
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		{
			name:     "retorna nil para proyecto nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja owner nil",
			input: &project.Project{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
				Owner:     nil,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expected: &ProjectResponse{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
				Owner:     nil,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToProjectResponse(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToOwnerResponse(t *testing.T) {
	userName := "John Doe"

	tests := []struct {
		name     string
		input    *user.User
		expected *OwnerResponse
	}{
		{
			name: "convierte user válido a owner response",
			input: &user.User{
				ID:    1,
				Name:  &userName,
				Email: "john@example.com",
			},
			expected: &OwnerResponse{
				ID:    1,
				Name:  &userName,
				Email: "john@example.com",
			},
		},
		{
			name:     "retorna nil para user nil",
			input:    nil,
			expected: nil,
		},
		{
			name: "maneja name nil",
			input: &user.User{
				ID:    1,
				Name:  nil,
				Email: "john@example.com",
			},
			expected: &OwnerResponse{
				ID:    1,
				Name:  nil,
				Email: "john@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToOwnerResponse(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToProjectListResponse(t *testing.T) {
	name1 := "Project 1"
	name2 := "Project 2"
	now := time.Now()

	tests := []struct {
		name     string
		input    []project.Project
		expected []ProjectResponse
	}{
		{
			name: "convierte lista de dominios a responses",
			input: []project.Project{
				{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now, UpdatedAt: now},
				{ID: 2, Name: &name2, CreatedBy: 2, CreatedAt: now, UpdatedAt: now},
			},
			expected: []ProjectResponse{
				{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now, UpdatedAt: now},
				{ID: 2, Name: &name2, CreatedBy: 2, CreatedAt: now, UpdatedAt: now},
			},
		},
		{
			name:     "maneja lista vacía",
			input:    []project.Project{},
			expected: []ProjectResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToProjectListResponse(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

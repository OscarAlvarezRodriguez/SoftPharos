package role

import (
	"context"
	"errors"
	"softpharos/internal/core/domain/role"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	desc1 := "Administrator"
	desc2 := "Developer"
	now := time.Now()

	tests := []struct {
		name          string
		mockSetup     func(*mockRepo.MockRoleRepository)
		expectedRoles []role.Role
		expectedErr   error
	}{
		{
			name: "retorna roles exitosamente",
			mockSetup: func(m *mockRepo.MockRoleRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]role.Role{
						{ID: 1, Name: "Admin", Description: &desc1, CreatedAt: now},
						{ID: 2, Name: "Developer", Description: &desc2, CreatedAt: now},
					}, nil)
			},
			expectedRoles: []role.Role{
				{ID: 1, Name: "Admin", Description: &desc1, CreatedAt: now},
				{ID: 2, Name: "Developer", Description: &desc2, CreatedAt: now},
			},
			expectedErr: nil,
		},
		{
			name: "retorna error cuando el repositorio falla",
			mockSetup: func(m *mockRepo.MockRoleRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]role.Role{}, errors.New("database error"))
			},
			expectedRoles: []role.Role{},
			expectedErr:   errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockRoleRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetAllRoles(ctx)

			assert.Equal(t, tt.expectedRoles, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetRoleByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	desc := "Administrator"
	now := time.Now()

	tests := []struct {
		name         string
		roleID       int
		mockSetup    func(*mockRepo.MockRoleRepository)
		expectedRole *role.Role
		expectedErr  error
	}{
		{
			name:   "retorna rol exitosamente",
			roleID: 1,
			mockSetup: func(m *mockRepo.MockRoleRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 1).
					Return(&role.Role{ID: 1, Name: "Admin", Description: &desc, CreatedAt: now}, nil)
			},
			expectedRole: &role.Role{ID: 1, Name: "Admin", Description: &desc, CreatedAt: now},
			expectedErr:  nil,
		},
		{
			name:   "retorna error cuando rol no existe",
			roleID: 999,
			mockSetup: func(m *mockRepo.MockRoleRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 999).
					Return(nil, errors.New("role not found"))
			},
			expectedRole: nil,
			expectedErr:  errors.New("role not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockRoleRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetRoleByID(ctx, tt.roleID)

			assert.Equal(t, tt.expectedRole, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetRoleByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	desc := "Administrator"
	now := time.Now()

	tests := []struct {
		name         string
		roleName     string
		mockSetup    func(*mockRepo.MockRoleRepository)
		expectedRole *role.Role
		expectedErr  error
	}{
		{
			name:     "retorna rol por nombre exitosamente",
			roleName: "Admin",
			mockSetup: func(m *mockRepo.MockRoleRepository) {
				m.EXPECT().
					GetByName(gomock.Any(), "Admin").
					Return(&role.Role{ID: 1, Name: "Admin", Description: &desc, CreatedAt: now}, nil)
			},
			expectedRole: &role.Role{ID: 1, Name: "Admin", Description: &desc, CreatedAt: now},
			expectedErr:  nil,
		},
		{
			name:     "retorna error cuando rol no existe",
			roleName: "NonExistent",
			mockSetup: func(m *mockRepo.MockRoleRepository) {
				m.EXPECT().
					GetByName(gomock.Any(), "NonExistent").
					Return(nil, errors.New("role not found"))
			},
			expectedRole: nil,
			expectedErr:  errors.New("role not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockRoleRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetRoleByName(ctx, tt.roleName)

			assert.Equal(t, tt.expectedRole, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

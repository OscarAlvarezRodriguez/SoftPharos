package user

import (
	"context"
	"errors"
	"softpharos/internal/core/domain/user"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name1 := "User 1"
	name2 := "User 2"
	now := time.Now()

	tests := []struct {
		name          string
		mockSetup     func(*mockRepo.MockUserRepository)
		expectedUsers []user.User
		expectedErr   error
	}{
		{
			name: "retorna usuarios exitosamente",
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]user.User{
						{ID: 1, Name: &name1, Email: "user1@example.com", RoleID: 1, CreatedAt: now},
						{ID: 2, Name: &name2, Email: "user2@example.com", RoleID: 2, CreatedAt: now},
					}, nil)
			},
			expectedUsers: []user.User{
				{ID: 1, Name: &name1, Email: "user1@example.com", RoleID: 1, CreatedAt: now},
				{ID: 2, Name: &name2, Email: "user2@example.com", RoleID: 2, CreatedAt: now},
			},
			expectedErr: nil,
		},
		{
			name: "retorna error cuando el repositorio falla",
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]user.User{}, errors.New("database error"))
			},
			expectedUsers: []user.User{},
			expectedErr:   errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockUserRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetAllUsers(ctx)

			assert.Equal(t, tt.expectedUsers, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Test User"
	now := time.Now()

	tests := []struct {
		name         string
		userID       int
		mockSetup    func(*mockRepo.MockUserRepository)
		expectedUser *user.User
		expectedErr  error
	}{
		{
			name:   "retorna usuario exitosamente",
			userID: 1,
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 1).
					Return(&user.User{ID: 1, Name: &name, Email: "test@example.com", RoleID: 1, CreatedAt: now}, nil)
			},
			expectedUser: &user.User{ID: 1, Name: &name, Email: "test@example.com", RoleID: 1, CreatedAt: now},
			expectedErr:  nil,
		},
		{
			name:   "retorna error cuando usuario no existe",
			userID: 999,
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 999).
					Return(nil, errors.New("user not found"))
			},
			expectedUser: nil,
			expectedErr:  errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockUserRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetUserByID(ctx, tt.userID)

			assert.Equal(t, tt.expectedUser, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Test User"
	email := "test@example.com"
	now := time.Now()

	tests := []struct {
		name         string
		email        string
		mockSetup    func(*mockRepo.MockUserRepository)
		expectedUser *user.User
		expectedErr  error
	}{
		{
			name:  "retorna usuario exitosamente",
			email: email,
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					GetByEmail(gomock.Any(), email).
					Return(&user.User{ID: 1, Name: &name, Email: email, RoleID: 1, CreatedAt: now}, nil)
			},
			expectedUser: &user.User{ID: 1, Name: &name, Email: email, RoleID: 1, CreatedAt: now},
			expectedErr:  nil,
		},
		{
			name:  "retorna error cuando email no existe",
			email: "notfound@example.com",
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					GetByEmail(gomock.Any(), "notfound@example.com").
					Return(nil, errors.New("user not found"))
			},
			expectedUser: nil,
			expectedErr:  errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockUserRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetUserByEmail(ctx, tt.email)

			assert.Equal(t, tt.expectedUser, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "New User"

	tests := []struct {
		name        string
		user        *user.User
		mockSetup   func(*mockRepo.MockUserRepository)
		expectedErr error
	}{
		{
			name: "crea usuario exitosamente",
			user: &user.User{Name: &name, Email: "new@example.com", Password: "hash", RoleID: 1},
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "retorna error cuando el repositorio falla",
			user: &user.User{Name: &name, Email: "new@example.com", Password: "hash", RoleID: 1},
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockUserRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.CreateUser(ctx, tt.user)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Updated User"

	tests := []struct {
		name        string
		user        *user.User
		mockSetup   func(*mockRepo.MockUserRepository)
		expectedErr error
	}{
		{
			name: "actualiza usuario exitosamente",
			user: &user.User{ID: 1, Name: &name, Email: "updated@example.com", Password: "hash", RoleID: 1},
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "retorna error cuando el repositorio falla",
			user: &user.User{ID: 1, Name: &name, Email: "updated@example.com", Password: "hash", RoleID: 1},
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockUserRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.UpdateUser(ctx, tt.user)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		userID      int
		mockSetup   func(*mockRepo.MockUserRepository)
		expectedErr error
	}{
		{
			name:   "elimina usuario exitosamente",
			userID: 1,
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					Delete(gomock.Any(), 1).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:   "retorna error cuando el repositorio falla",
			userID: 1,
			mockSetup: func(m *mockRepo.MockUserRepository) {
				m.EXPECT().
					Delete(gomock.Any(), 1).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockUserRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.DeleteUser(ctx, tt.userID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

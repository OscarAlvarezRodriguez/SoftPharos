package project

import (
	"context"
	"errors"
	"softpharos/internal/core/domain/project"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllProjects(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name1 := "Project 1"
	name2 := "Project 2"
	now := time.Now()

	tests := []struct {
		name          string
		mockSetup     func(*mockRepo.MockProjectRepository)
		expectedProjs []project.Project
		expectedErr   error
	}{
		{
			name: "retorna proyectos exitosamente",
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]project.Project{
						{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now},
						{ID: 2, Name: &name2, CreatedBy: 2, CreatedAt: now},
					}, nil)
			},
			expectedProjs: []project.Project{
				{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now},
				{ID: 2, Name: &name2, CreatedBy: 2, CreatedAt: now},
			},
			expectedErr: nil,
		},
		{
			name: "retorna error cuando el repositorio falla",
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]project.Project{}, errors.New("database error"))
			},
			expectedProjs: []project.Project{},
			expectedErr:   errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockProjectRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetAllProjects(ctx)

			assert.Equal(t, tt.expectedProjs, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetProjectByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Test Project"
	now := time.Now()

	tests := []struct {
		name         string
		projectID    int
		mockSetup    func(*mockRepo.MockProjectRepository)
		expectedProj *project.Project
		expectedErr  error
	}{
		{
			name:      "retorna proyecto exitosamente",
			projectID: 1,
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 1).
					Return(&project.Project{ID: 1, Name: &name, CreatedBy: 1, CreatedAt: now}, nil)
			},
			expectedProj: &project.Project{ID: 1, Name: &name, CreatedBy: 1, CreatedAt: now},
			expectedErr:  nil,
		},
		{
			name:      "retorna error cuando proyecto no existe",
			projectID: 999,
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 999).
					Return(nil, errors.New("project not found"))
			},
			expectedProj: nil,
			expectedErr:  errors.New("project not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockProjectRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetProjectByID(ctx, tt.projectID)

			assert.Equal(t, tt.expectedProj, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetProjectsByOwner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name1 := "Project 1"
	name2 := "Project 2"
	now := time.Now()

	tests := []struct {
		name          string
		ownerID       int
		mockSetup     func(*mockRepo.MockProjectRepository)
		expectedProjs []project.Project
		expectedErr   error
	}{
		{
			name:    "retorna proyectos del owner exitosamente",
			ownerID: 1,
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					GetByOwner(gomock.Any(), 1).
					Return([]project.Project{
						{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now},
						{ID: 2, Name: &name2, CreatedBy: 1, CreatedAt: now},
					}, nil)
			},
			expectedProjs: []project.Project{
				{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now},
				{ID: 2, Name: &name2, CreatedBy: 1, CreatedAt: now},
			},
			expectedErr: nil,
		},
		{
			name:    "retorna lista vacía cuando owner no tiene proyectos",
			ownerID: 999,
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					GetByOwner(gomock.Any(), 999).
					Return([]project.Project{}, nil)
			},
			expectedProjs: []project.Project{},
			expectedErr:   nil,
		},
		{
			name:    "retorna error cuando el repositorio falla",
			ownerID: 1,
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					GetByOwner(gomock.Any(), 1).
					Return([]project.Project{}, errors.New("database error"))
			},
			expectedProjs: []project.Project{},
			expectedErr:   errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockProjectRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetProjectsByOwner(ctx, tt.ownerID)

			assert.Equal(t, tt.expectedProjs, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "New Project"

	tests := []struct {
		name        string
		project     *project.Project
		mockSetup   func(*mockRepo.MockProjectRepository)
		expectedErr error
	}{
		{
			name:    "crea proyecto exitosamente",
			project: &project.Project{Name: &name, CreatedBy: 1},
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:    "retorna error cuando el repositorio falla",
			project: &project.Project{Name: &name, CreatedBy: 1},
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockProjectRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.CreateProject(ctx, tt.project)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Updated Project"

	tests := []struct {
		name        string
		project     *project.Project
		mockSetup   func(*mockRepo.MockProjectRepository)
		expectedErr error
	}{
		{
			name:    "actualiza proyecto exitosamente",
			project: &project.Project{ID: 1, Name: &name, CreatedBy: 1},
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:    "retorna error cuando el repositorio falla",
			project: &project.Project{ID: 1, Name: &name, CreatedBy: 1},
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockProjectRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.UpdateProject(ctx, tt.project)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		projectID   int
		mockSetup   func(*mockRepo.MockProjectRepository)
		expectedErr error
	}{
		{
			name:      "elimina proyecto exitosamente",
			projectID: 1,
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					Delete(gomock.Any(), 1).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:      "retorna error cuando el repositorio falla",
			projectID: 1,
			mockSetup: func(m *mockRepo.MockProjectRepository) {
				m.EXPECT().
					Delete(gomock.Any(), 1).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockProjectRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.DeleteProject(ctx, tt.projectID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

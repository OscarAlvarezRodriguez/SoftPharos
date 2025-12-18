package milestone

import (
	"context"
	"errors"
	"softpharos/internal/core/domain/milestone"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllMilestones(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title1 := "Milestone 1"
	title2 := "Milestone 2"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockRepo.MockMilestoneRepository)
		expectedMilestones []milestone.Milestone
		expectedErr        error
	}{
		{
			name: "retorna milestones exitosamente",
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]milestone.Milestone{
						{ID: 1, ProjectID: 1, Title: &title1, CreatedAt: now},
						{ID: 2, ProjectID: 1, Title: &title2, CreatedAt: now},
					}, nil)
			},
			expectedMilestones: []milestone.Milestone{
				{ID: 1, ProjectID: 1, Title: &title1, CreatedAt: now},
				{ID: 2, ProjectID: 1, Title: &title2, CreatedAt: now},
			},
			expectedErr: nil,
		},
		{
			name: "retorna error cuando el repositorio falla",
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					GetAll(gomock.Any()).
					Return([]milestone.Milestone{}, errors.New("database error"))
			},
			expectedMilestones: []milestone.Milestone{},
			expectedErr:        errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockMilestoneRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetAllMilestones(ctx)

			assert.Equal(t, tt.expectedMilestones, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetMilestoneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "Test Milestone"
	now := time.Now()

	tests := []struct {
		name              string
		milestoneID       int
		mockSetup         func(*mockRepo.MockMilestoneRepository)
		expectedMilestone *milestone.Milestone
		expectedErr       error
	}{
		{
			name:        "retorna milestone exitosamente",
			milestoneID: 1,
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 1).
					Return(&milestone.Milestone{ID: 1, ProjectID: 1, Title: &title, CreatedAt: now}, nil)
			},
			expectedMilestone: &milestone.Milestone{ID: 1, ProjectID: 1, Title: &title, CreatedAt: now},
			expectedErr:       nil,
		},
		{
			name:        "retorna error cuando milestone no existe",
			milestoneID: 999,
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					GetByID(gomock.Any(), 999).
					Return(nil, errors.New("milestone not found"))
			},
			expectedMilestone: nil,
			expectedErr:       errors.New("milestone not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockMilestoneRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetMilestoneByID(ctx, tt.milestoneID)

			assert.Equal(t, tt.expectedMilestone, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetMilestonesByProjectID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title1 := "Milestone 1"
	title2 := "Milestone 2"
	now := time.Now()

	tests := []struct {
		name               string
		projectID          int
		mockSetup          func(*mockRepo.MockMilestoneRepository)
		expectedMilestones []milestone.Milestone
		expectedErr        error
	}{
		{
			name:      "retorna milestones del proyecto exitosamente",
			projectID: 1,
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					GetByProjectID(gomock.Any(), 1).
					Return([]milestone.Milestone{
						{ID: 1, ProjectID: 1, Title: &title1, CreatedAt: now},
						{ID: 2, ProjectID: 1, Title: &title2, CreatedAt: now},
					}, nil)
			},
			expectedMilestones: []milestone.Milestone{
				{ID: 1, ProjectID: 1, Title: &title1, CreatedAt: now},
				{ID: 2, ProjectID: 1, Title: &title2, CreatedAt: now},
			},
			expectedErr: nil,
		},
		{
			name:      "retorna error cuando el repositorio falla",
			projectID: 1,
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					GetByProjectID(gomock.Any(), 1).
					Return([]milestone.Milestone{}, errors.New("database error"))
			},
			expectedMilestones: []milestone.Milestone{},
			expectedErr:        errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockMilestoneRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			result, err := service.GetMilestonesByProjectID(ctx, tt.projectID)

			assert.Equal(t, tt.expectedMilestones, result)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateMilestone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "New Milestone"

	tests := []struct {
		name        string
		milestone   *milestone.Milestone
		mockSetup   func(*mockRepo.MockMilestoneRepository)
		expectedErr error
	}{
		{
			name:      "crea milestone exitosamente",
			milestone: &milestone.Milestone{ProjectID: 1, Title: &title},
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:      "retorna error cuando el repositorio falla",
			milestone: &milestone.Milestone{ProjectID: 1, Title: &title},
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockMilestoneRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.CreateMilestone(ctx, tt.milestone)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateMilestone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "Updated Milestone"

	tests := []struct {
		name        string
		milestone   *milestone.Milestone
		mockSetup   func(*mockRepo.MockMilestoneRepository)
		expectedErr error
	}{
		{
			name:      "actualiza milestone exitosamente",
			milestone: &milestone.Milestone{ID: 1, ProjectID: 1, Title: &title},
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:      "retorna error cuando el repositorio falla",
			milestone: &milestone.Milestone{ID: 1, ProjectID: 1, Title: &title},
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockMilestoneRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.UpdateMilestone(ctx, tt.milestone)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteMilestone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		milestoneID int
		mockSetup   func(*mockRepo.MockMilestoneRepository)
		expectedErr error
	}{
		{
			name:        "elimina milestone exitosamente",
			milestoneID: 1,
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					Delete(gomock.Any(), 1).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:        "retorna error cuando el repositorio falla",
			milestoneID: 1,
			mockSetup: func(m *mockRepo.MockMilestoneRepository) {
				m.EXPECT().
					Delete(gomock.Any(), 1).
					Return(errors.New("database error"))
			},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mockRepo.NewMockMilestoneRepository(ctrl)
			tt.mockSetup(mockRepository)

			service := New(mockRepository)
			ctx := context.Background()

			err := service.DeleteMilestone(ctx, tt.milestoneID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

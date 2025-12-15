package feedback

import (
	"context"
	"softpharos/internal/core/domain/feedback"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllFeedbacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	mockRepo := mockRepo.NewMockFeedbackRepository(ctrl)
	mockRepo.EXPECT().GetAll(gomock.Any()).Return([]feedback.Feedback{
		{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Good work", CreatedAt: now},
	}, nil)

	service := New(mockRepo)
	result, err := service.GetAllFeedbacks(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestGetFeedbackByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	mockRepo := mockRepo.NewMockFeedbackRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), 1).Return(&feedback.Feedback{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Good work", CreatedAt: now}, nil)

	service := New(mockRepo)
	result, err := service.GetFeedbackByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetFeedbacksByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockFeedbackRepository(ctrl)
	mockRepo.EXPECT().GetByMilestoneID(gomock.Any(), 1).Return([]feedback.Feedback{{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Good"}}, nil)

	service := New(mockRepo)
	result, err := service.GetFeedbacksByMilestoneID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestCreateFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockFeedbackRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.CreateFeedback(context.Background(), &feedback.Feedback{MilestoneID: 1, ProfessorID: 1, Content: "Good"})

	assert.NoError(t, err)
}

func TestUpdateFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockFeedbackRepository(ctrl)
	mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.UpdateFeedback(context.Background(), &feedback.Feedback{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Good"})

	assert.NoError(t, err)
}

func TestDeleteFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockFeedbackRepository(ctrl)
	mockRepo.EXPECT().Delete(gomock.Any(), 1).Return(nil)

	service := New(mockRepo)
	err := service.DeleteFeedback(context.Background(), 1)

	assert.NoError(t, err)
}

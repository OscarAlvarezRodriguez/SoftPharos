package deliverable

import (
	"context"
	"softpharos/internal/core/domain/deliverable"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllDeliverables(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	typeVal := "PDF"
	now := time.Now()

	mockRepo := mockRepo.NewMockDeliverableRepository(ctrl)
	mockRepo.EXPECT().GetAll(gomock.Any()).Return([]deliverable.Deliverable{
		{ID: 1, MilestoneID: 1, URL: "http://example.com", Type: &typeVal, CreatedAt: now},
	}, nil)

	service := New(mockRepo)
	result, err := service.GetAllDeliverables(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestGetDeliverableByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	typeVal := "PDF"
	now := time.Now()

	mockRepo := mockRepo.NewMockDeliverableRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), 1).Return(&deliverable.Deliverable{ID: 1, MilestoneID: 1, URL: "http://example.com", Type: &typeVal, CreatedAt: now}, nil)

	service := New(mockRepo)
	result, err := service.GetDeliverableByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetDeliverablesByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockDeliverableRepository(ctrl)
	mockRepo.EXPECT().GetByMilestoneID(gomock.Any(), 1).Return([]deliverable.Deliverable{{ID: 1, MilestoneID: 1, URL: "http://example.com"}}, nil)

	service := New(mockRepo)
	result, err := service.GetDeliverablesByMilestoneID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestCreateDeliverable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockDeliverableRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.CreateDeliverable(context.Background(), &deliverable.Deliverable{MilestoneID: 1, URL: "http://example.com"})

	assert.NoError(t, err)
}

func TestUpdateDeliverable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockDeliverableRepository(ctrl)
	mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.UpdateDeliverable(context.Background(), &deliverable.Deliverable{ID: 1, MilestoneID: 1, URL: "http://example.com"})

	assert.NoError(t, err)
}

func TestDeleteDeliverable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockDeliverableRepository(ctrl)
	mockRepo.EXPECT().Delete(gomock.Any(), 1).Return(nil)

	service := New(mockRepo)
	err := service.DeleteDeliverable(context.Background(), 1)

	assert.NoError(t, err)
}

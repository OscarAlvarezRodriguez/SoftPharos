package reaction

import (
	"context"
	"softpharos/internal/core/domain/reaction"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllReactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reactionType := "like"
	now := time.Now()

	mockRepo := mockRepo.NewMockReactionRepository(ctrl)
	mockRepo.EXPECT().GetAll(gomock.Any()).Return([]reaction.Reaction{
		{ID: 1, MilestoneID: 1, UserID: 1, Type: &reactionType, CreatedAt: now},
	}, nil)

	service := New(mockRepo)
	result, err := service.GetAllReactions(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestGetReactionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reactionType := "like"
	now := time.Now()

	mockRepo := mockRepo.NewMockReactionRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), 1).Return(&reaction.Reaction{ID: 1, MilestoneID: 1, UserID: 1, Type: &reactionType, CreatedAt: now}, nil)

	service := New(mockRepo)
	result, err := service.GetReactionByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetReactionsByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockReactionRepository(ctrl)
	mockRepo.EXPECT().GetByMilestoneID(gomock.Any(), 1).Return([]reaction.Reaction{{ID: 1, MilestoneID: 1, UserID: 1}}, nil)

	service := New(mockRepo)
	result, err := service.GetReactionsByMilestoneID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestCreateReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockReactionRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.CreateReaction(context.Background(), &reaction.Reaction{MilestoneID: 1, UserID: 1})

	assert.NoError(t, err)
}

func TestUpdateReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockReactionRepository(ctrl)
	mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.UpdateReaction(context.Background(), &reaction.Reaction{ID: 1, MilestoneID: 1, UserID: 1})

	assert.NoError(t, err)
}

func TestDeleteReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockReactionRepository(ctrl)
	mockRepo.EXPECT().Delete(gomock.Any(), 1).Return(nil)

	service := New(mockRepo)
	err := service.DeleteReaction(context.Background(), 1)

	assert.NoError(t, err)
}

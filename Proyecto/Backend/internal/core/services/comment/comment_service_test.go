package comment

import (
	"context"
	"softpharos/internal/core/domain/comment"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content1 := "Comment 1"
	content2 := "Comment 2"
	now := time.Now()

	mockRepo := mockRepo.NewMockCommentRepository(ctrl)
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return([]comment.Comment{
			{ID: 1, MilestoneID: 1, UserID: 1, Content: &content1, CreatedAt: now},
			{ID: 2, MilestoneID: 1, UserID: 1, Content: &content2, CreatedAt: now},
		}, nil)

	service := New(mockRepo)
	result, err := service.GetAllComments(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestGetCommentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content := "Test Comment"
	now := time.Now()

	mockRepo := mockRepo.NewMockCommentRepository(ctrl)
	mockRepo.EXPECT().
		GetByID(gomock.Any(), 1).
		Return(&comment.Comment{ID: 1, MilestoneID: 1, UserID: 1, Content: &content, CreatedAt: now}, nil)

	service := New(mockRepo)
	result, err := service.GetCommentByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

func TestGetCommentsByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content1 := "Comment 1"
	now := time.Now()

	mockRepo := mockRepo.NewMockCommentRepository(ctrl)
	mockRepo.EXPECT().
		GetByMilestoneID(gomock.Any(), 1).
		Return([]comment.Comment{
			{ID: 1, MilestoneID: 1, UserID: 1, Content: &content1, CreatedAt: now},
		}, nil)

	service := New(mockRepo)
	result, err := service.GetCommentsByMilestoneID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content := "New Comment"

	mockRepo := mockRepo.NewMockCommentRepository(ctrl)
	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	service := New(mockRepo)
	err := service.CreateComment(context.Background(), &comment.Comment{MilestoneID: 1, UserID: 1, Content: &content})

	assert.NoError(t, err)
}

func TestUpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content := "Updated Comment"

	mockRepo := mockRepo.NewMockCommentRepository(ctrl)
	mockRepo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	service := New(mockRepo)
	err := service.UpdateComment(context.Background(), &comment.Comment{ID: 1, MilestoneID: 1, UserID: 1, Content: &content})

	assert.NoError(t, err)
}

func TestDeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockCommentRepository(ctrl)
	mockRepo.EXPECT().
		Delete(gomock.Any(), 1).
		Return(nil)

	service := New(mockRepo)
	err := service.DeleteComment(context.Background(), 1)

	assert.NoError(t, err)
}

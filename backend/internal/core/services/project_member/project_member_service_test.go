package project_member

import (
	"context"
	"softpharos/internal/core/domain/project_member"
	mockRepo "softpharos/mocks/core/ports/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllProjectMembers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := "Developer"
	now := time.Now()

	mockRepo := mockRepo.NewMockProjectMemberRepository(ctrl)
	mockRepo.EXPECT().GetAll(gomock.Any()).Return([]project_member.ProjectMember{
		{ID: 1, ProjectID: 1, UserID: 1, Role: &role, JoinedAt: now},
	}, nil)

	service := New(mockRepo)
	result, err := service.GetAllProjectMembers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestGetProjectMemberByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := "Developer"
	now := time.Now()

	mockRepo := mockRepo.NewMockProjectMemberRepository(ctrl)
	mockRepo.EXPECT().GetByID(gomock.Any(), 1).Return(&project_member.ProjectMember{ID: 1, ProjectID: 1, UserID: 1, Role: &role, JoinedAt: now}, nil)

	service := New(mockRepo)
	result, err := service.GetProjectMemberByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetProjectMembersByProjectID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockProjectMemberRepository(ctrl)
	mockRepo.EXPECT().GetByProjectID(gomock.Any(), 1).Return([]project_member.ProjectMember{{ID: 1, ProjectID: 1, UserID: 1}}, nil)

	service := New(mockRepo)
	result, err := service.GetProjectMembersByProjectID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestCreateProjectMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockProjectMemberRepository(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.CreateProjectMember(context.Background(), &project_member.ProjectMember{ProjectID: 1, UserID: 1})

	assert.NoError(t, err)
}

func TestUpdateProjectMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockProjectMemberRepository(ctrl)
	mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

	service := New(mockRepo)
	err := service.UpdateProjectMember(context.Background(), &project_member.ProjectMember{ID: 1, ProjectID: 1, UserID: 1})

	assert.NoError(t, err)
}

func TestDeleteProjectMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockRepo.NewMockProjectMemberRepository(ctrl)
	mockRepo.EXPECT().Delete(gomock.Any(), 1).Return(nil)

	service := New(mockRepo)
	err := service.DeleteProjectMember(context.Background(), 1)

	assert.NoError(t, err)
}

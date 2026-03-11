package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mockRepo "softpharos/mocks/core/ports/repository"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRepo.NewMockUserRepository(ctrl)
	mockRoleRepo := mockRepo.NewMockRoleRepository(ctrl)

	service := New(mockUserRepo, mockRoleRepo)

	assert.NotNil(t, service)

	// Verificar que el service tiene los repositorios correctos
	svc, ok := service.(*Service)
	assert.True(t, ok)
	assert.NotNil(t, svc.userRepo)
	assert.NotNil(t, svc.roleRepo)
}

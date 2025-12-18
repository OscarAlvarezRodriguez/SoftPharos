package role

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"softpharos/internal/core/domain/role"
	mockService "softpharos/mocks/core/ports/services"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	desc1 := "Administrator"
	desc2 := "Developer"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockRoleService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los roles exitosamente",
			mockSetup: func(m *mockService.MockRoleService) {
				m.EXPECT().
					GetAllRoles(gomock.Any()).
					Return([]role.Role{
						{ID: 1, Name: "Admin", Description: &desc1, CreatedAt: now},
						{ID: 2, Name: "Developer", Description: &desc2, CreatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockRoleService) {
				m.EXPECT().
					GetAllRoles(gomock.Any()).
					Return([]role.Role{}, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockRoleService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/roles", controller.GetAllRoles)

			req, _ := http.NewRequest("GET", "/roles", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetRoleByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	desc := "Administrator"
	now := time.Now()

	tests := []struct {
		name               string
		roleID             string
		mockSetup          func(*mockService.MockRoleService)
		expectedStatusCode int
	}{
		{
			name:   "retorna rol exitosamente",
			roleID: "1",
			mockSetup: func(m *mockService.MockRoleService) {
				m.EXPECT().
					GetRoleByID(gomock.Any(), 1).
					Return(&role.Role{
						ID:          1,
						Name:        "Admin",
						Description: &desc,
						CreatedAt:   now,
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:   "retorna error cuando el rol no existe",
			roleID: "999",
			mockSetup: func(m *mockService.MockRoleService) {
				m.EXPECT().
					GetRoleByID(gomock.Any(), 999).
					Return(nil, errors.New("role not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "retorna error cuando el ID es inválido",
			roleID:             "invalid",
			mockSetup:          func(m *mockService.MockRoleService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockRoleService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/roles/:id", controller.GetRoleByID)

			req, _ := http.NewRequest("GET", "/roles/"+tt.roleID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetRoleByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	desc := "Administrator"
	now := time.Now()

	tests := []struct {
		name               string
		roleName           string
		mockSetup          func(*mockService.MockRoleService)
		expectedStatusCode int
	}{
		{
			name:     "retorna rol por nombre exitosamente",
			roleName: "Admin",
			mockSetup: func(m *mockService.MockRoleService) {
				m.EXPECT().
					GetRoleByName(gomock.Any(), "Admin").
					Return(&role.Role{
						ID:          1,
						Name:        "Admin",
						Description: &desc,
						CreatedAt:   now,
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:     "retorna error cuando el rol no existe",
			roleName: "NonExistent",
			mockSetup: func(m *mockService.MockRoleService) {
				m.EXPECT().
					GetRoleByName(gomock.Any(), "NonExistent").
					Return(nil, errors.New("role not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockRoleService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/roles/name/:name", controller.GetRoleByName)

			req, _ := http.NewRequest("GET", "/roles/name/"+tt.roleName, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

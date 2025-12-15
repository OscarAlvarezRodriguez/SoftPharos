package project_member

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"softpharos/internal/core/domain/project_member"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllProjectMembers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role1 := "Developer"
	role2 := "Designer"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockProjectMemberService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los miembros exitosamente",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetAllProjectMembers(gomock.Any()).Return([]project_member.ProjectMember{
					{ID: 1, ProjectID: 1, UserID: 1, Role: &role1, JoinedAt: now},
					{ID: 2, ProjectID: 1, UserID: 2, Role: &role2, JoinedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetAllProjectMembers(gomock.Any()).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectMemberService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/project-members", controller.GetAllProjectMembers)
			req, _ := http.NewRequest("GET", "/project-members", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetProjectMemberByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := "Developer"
	now := time.Now()

	tests := []struct {
		name               string
		memberID           string
		mockSetup          func(*mockService.MockProjectMemberService)
		expectedStatusCode int
	}{
		{
			name:     "retorna miembro exitosamente",
			memberID: "1",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetProjectMemberByID(gomock.Any(), 1).Return(&project_member.ProjectMember{
					ID: 1, ProjectID: 1, UserID: 1, Role: &role, JoinedAt: now,
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			memberID:           "invalid",
			mockSetup:          func(m *mockService.MockProjectMemberService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:     "retorna error cuando miembro no existe",
			memberID: "999",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetProjectMemberByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectMemberService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/project-members/:id", controller.GetProjectMemberByID)
			req, _ := http.NewRequest("GET", "/project-members/"+tt.memberID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetProjectMembersByProjectID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := "Developer"
	now := time.Now()

	tests := []struct {
		name               string
		projectID          string
		mockSetup          func(*mockService.MockProjectMemberService)
		expectedStatusCode int
	}{
		{
			name:      "retorna miembros por projectID exitosamente",
			projectID: "1",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetProjectMembersByProjectID(gomock.Any(), 1).Return([]project_member.ProjectMember{
					{ID: 1, ProjectID: 1, UserID: 1, Role: &role, JoinedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para projectID inválido",
			projectID:          "invalid",
			mockSetup:          func(m *mockService.MockProjectMemberService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando el service falla",
			projectID: "1",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetProjectMembersByProjectID(gomock.Any(), 1).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectMemberService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/project-members/project/:projectId", controller.GetProjectMembersByProjectID)
			req, _ := http.NewRequest("GET", "/project-members/project/"+tt.projectID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateProjectMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := "Developer"

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockProjectMemberService)
		expectedStatusCode int
	}{
		{
			name:        "crea miembro exitosamente",
			requestBody: CreateProjectMemberRequest{ProjectID: 1, UserID: 1, Role: &role},
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().CreateProjectMember(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockProjectMemberService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			requestBody: CreateProjectMemberRequest{ProjectID: 1, UserID: 1, Role: &role},
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().CreateProjectMember(gomock.Any(), gomock.Any()).Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectMemberService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/project-members", controller.CreateProjectMember)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("POST", "/project-members", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateProjectMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := "Developer"
	updatedRole := "Lead Developer"
	now := time.Now()

	tests := []struct {
		name               string
		memberID           string
		requestBody        interface{}
		mockSetup          func(*mockService.MockProjectMemberService)
		expectedStatusCode int
	}{
		{
			name:        "actualiza miembro exitosamente",
			memberID:    "1",
			requestBody: UpdateProjectMemberRequest{Role: &updatedRole},
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetProjectMemberByID(gomock.Any(), 1).Return(&project_member.ProjectMember{
					ID: 1, ProjectID: 1, UserID: 1, Role: &role, JoinedAt: now,
				}, nil)
				m.EXPECT().UpdateProjectMember(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			memberID:           "invalid",
			requestBody:        UpdateProjectMemberRequest{},
			mockSetup:          func(m *mockService.MockProjectMemberService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando miembro no existe",
			memberID:    "999",
			requestBody: UpdateProjectMemberRequest{Role: &updatedRole},
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetProjectMemberByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:        "retorna error cuando update falla",
			memberID:    "1",
			requestBody: UpdateProjectMemberRequest{Role: &updatedRole},
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().GetProjectMemberByID(gomock.Any(), 1).Return(&project_member.ProjectMember{
					ID: 1, ProjectID: 1, UserID: 1, Role: &role, JoinedAt: now,
				}, nil)
				m.EXPECT().UpdateProjectMember(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectMemberService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/project-members/:id", controller.UpdateProjectMember)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("PUT", "/project-members/"+tt.memberID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteProjectMember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		memberID           string
		mockSetup          func(*mockService.MockProjectMemberService)
		expectedStatusCode int
	}{
		{
			name:     "elimina miembro exitosamente",
			memberID: "1",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().DeleteProjectMember(gomock.Any(), 1).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			memberID:           "invalid",
			mockSetup:          func(m *mockService.MockProjectMemberService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:     "retorna error cuando delete falla",
			memberID: "1",
			mockSetup: func(m *mockService.MockProjectMemberService) {
				m.EXPECT().DeleteProjectMember(gomock.Any(), 1).Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectMemberService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/project-members/:id", controller.DeleteProjectMember)
			req, _ := http.NewRequest("DELETE", "/project-members/"+tt.memberID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

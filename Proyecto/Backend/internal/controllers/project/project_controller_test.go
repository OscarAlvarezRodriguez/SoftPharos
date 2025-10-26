package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/domain/user"
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

func TestGetAllProjects(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name1 := "Project 1"
	name2 := "Project 2"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockProjectService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los proyectos exitosamente",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetAllProjects(gomock.Any()).
					Return([]project.Project{
						{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now, UpdatedAt: now},
						{ID: 2, Name: &name2, CreatedBy: 2, CreatedAt: now, UpdatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetAllProjects(gomock.Any()).
					Return([]project.Project{}, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/projects", controller.GetAllProjects)

			req, _ := http.NewRequest("GET", "/projects", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetProjectByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Test Project"
	userName := "John Doe"
	now := time.Now()

	tests := []struct {
		name               string
		projectID          string
		mockSetup          func(*mockService.MockProjectService)
		expectedStatusCode int
	}{
		{
			name:      "retorna proyecto exitosamente",
			projectID: "1",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetProjectByID(gomock.Any(), 1).
					Return(&project.Project{
						ID:        1,
						Name:      &name,
						CreatedBy: 1,
						Owner: &user.User{
							ID:    1,
							Name:  &userName,
							Email: "john@example.com",
						},
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			projectID:          "invalid",
			mockSetup:          func(m *mockService.MockProjectService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando proyecto no existe",
			projectID: "999",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetProjectByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/projects/:id", controller.GetProjectByID)

			req, _ := http.NewRequest("GET", "/projects/"+tt.projectID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
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
		name               string
		ownerID            string
		mockSetup          func(*mockService.MockProjectService)
		expectedStatusCode int
	}{
		{
			name:    "retorna proyectos del owner exitosamente",
			ownerID: "1",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetProjectsByOwner(gomock.Any(), 1).
					Return([]project.Project{
						{ID: 1, Name: &name1, CreatedBy: 1, CreatedAt: now, UpdatedAt: now},
						{ID: 2, Name: &name2, CreatedBy: 1, CreatedAt: now, UpdatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para owner ID inválido",
			ownerID:            "invalid",
			mockSetup:          func(m *mockService.MockProjectService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "retorna error cuando el service falla",
			ownerID: "1",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetProjectsByOwner(gomock.Any(), 1).
					Return([]project.Project{}, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/projects/owner/:ownerId", controller.GetProjectsByOwner)

			req, _ := http.NewRequest("GET", "/projects/owner/"+tt.ownerID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "New Project"
	objective := "Test Objective"

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockProjectService)
		expectedStatusCode int
	}{
		{
			name: "crea proyecto exitosamente",
			requestBody: CreateProjectRequest{
				Name:      &name,
				Objective: &objective,
				CreatedBy: 1,
			},
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					CreateProject(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockProjectService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "retorna error cuando el service falla",
			requestBody: CreateProjectRequest{
				Name:      &name,
				CreatedBy: 1,
			},
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					CreateProject(gomock.Any(), gomock.Any()).
					Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/projects", controller.CreateProject)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("POST", "/projects", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Existing Project"
	updatedName := "Updated Name"
	now := time.Now()

	tests := []struct {
		name               string
		projectID          string
		requestBody        interface{}
		mockSetup          func(*mockService.MockProjectService)
		expectedStatusCode int
	}{
		{
			name:      "actualiza proyecto exitosamente",
			projectID: "1",
			requestBody: UpdateProjectRequest{
				Name: &updatedName,
			},
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetProjectByID(gomock.Any(), 1).
					Return(&project.Project{
						ID:        1,
						Name:      &name,
						CreatedBy: 1,
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
				m.EXPECT().
					UpdateProject(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			projectID:          "invalid",
			requestBody:        UpdateProjectRequest{},
			mockSetup:          func(m *mockService.MockProjectService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando proyecto no existe",
			projectID: "999",
			requestBody: UpdateProjectRequest{
				Name: &updatedName,
			},
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetProjectByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:      "retorna error cuando update falla",
			projectID: "1",
			requestBody: UpdateProjectRequest{
				Name: &updatedName,
			},
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					GetProjectByID(gomock.Any(), 1).
					Return(&project.Project{
						ID:        1,
						Name:      &name,
						CreatedBy: 1,
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
				m.EXPECT().
					UpdateProject(gomock.Any(), gomock.Any()).
					Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/projects/:id", controller.UpdateProject)

			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("PUT", "/projects/"+tt.projectID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		projectID          string
		mockSetup          func(*mockService.MockProjectService)
		expectedStatusCode int
	}{
		{
			name:      "elimina proyecto exitosamente",
			projectID: "1",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					DeleteProject(gomock.Any(), 1).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			projectID:          "invalid",
			mockSetup:          func(m *mockService.MockProjectService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando delete falla",
			projectID: "1",
			mockSetup: func(m *mockService.MockProjectService) {
				m.EXPECT().
					DeleteProject(gomock.Any(), 1).
					Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockProjectService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/projects/:id", controller.DeleteProject)

			req, _ := http.NewRequest("DELETE", "/projects/"+tt.projectID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

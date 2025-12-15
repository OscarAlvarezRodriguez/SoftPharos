package milestone

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

	"softpharos/internal/core/domain/milestone"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllMilestones(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title1 := "Milestone 1"
	title2 := "Milestone 2"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockMilestoneService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los milestones exitosamente",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetAllMilestones(gomock.Any()).
					Return([]milestone.Milestone{
						{ID: 1, ProjectID: 1, Title: &title1, CreatedAt: now},
						{ID: 2, ProjectID: 1, Title: &title2, CreatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetAllMilestones(gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockMilestoneService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/milestones", controller.GetAllMilestones)

			req, _ := http.NewRequest("GET", "/milestones", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetMilestoneByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "Test Milestone"
	now := time.Now()

	tests := []struct {
		name               string
		milestoneID        string
		mockSetup          func(*mockService.MockMilestoneService)
		expectedStatusCode int
	}{
		{
			name:        "retorna milestone exitosamente",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetMilestoneByID(gomock.Any(), 1).
					Return(&milestone.Milestone{
						ID:        1,
						ProjectID: 1,
						Title:     &title,
						CreatedAt: now,
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			milestoneID:        "invalid",
			mockSetup:          func(m *mockService.MockMilestoneService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando milestone no existe",
			milestoneID: "999",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetMilestoneByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockMilestoneService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/milestones/:id", controller.GetMilestoneByID)

			req, _ := http.NewRequest("GET", "/milestones/"+tt.milestoneID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
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
		projectID          string
		mockSetup          func(*mockService.MockMilestoneService)
		expectedStatusCode int
	}{
		{
			name:      "retorna milestones por projectID exitosamente",
			projectID: "1",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetMilestonesByProjectID(gomock.Any(), 1).
					Return([]milestone.Milestone{
						{ID: 1, ProjectID: 1, Title: &title1, CreatedAt: now},
						{ID: 2, ProjectID: 1, Title: &title2, CreatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para projectID inválido",
			projectID:          "invalid",
			mockSetup:          func(m *mockService.MockMilestoneService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando el service falla",
			projectID: "1",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetMilestonesByProjectID(gomock.Any(), 1).
					Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockMilestoneService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/milestones/project/:projectId", controller.GetMilestonesByProjectID)

			req, _ := http.NewRequest("GET", "/milestones/project/"+tt.projectID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateMilestone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "New Milestone"

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockMilestoneService)
		expectedStatusCode int
	}{
		{
			name: "crea milestone exitosamente",
			requestBody: CreateMilestoneRequest{
				ProjectID: 1,
				Title:     &title,
			},
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					CreateMilestone(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockMilestoneService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "retorna error cuando el service falla",
			requestBody: CreateMilestoneRequest{
				ProjectID: 1,
				Title:     &title,
			},
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					CreateMilestone(gomock.Any(), gomock.Any()).
					Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockMilestoneService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/milestones", controller.CreateMilestone)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("POST", "/milestones", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateMilestone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	title := "Existing Milestone"
	updatedTitle := "Updated Milestone"
	now := time.Now()

	tests := []struct {
		name               string
		milestoneID        string
		requestBody        interface{}
		mockSetup          func(*mockService.MockMilestoneService)
		expectedStatusCode int
	}{
		{
			name:        "actualiza milestone exitosamente",
			milestoneID: "1",
			requestBody: UpdateMilestoneRequest{
				Title: &updatedTitle,
			},
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetMilestoneByID(gomock.Any(), 1).
					Return(&milestone.Milestone{
						ID:        1,
						ProjectID: 1,
						Title:     &title,
						CreatedAt: now,
					}, nil)
				m.EXPECT().
					UpdateMilestone(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			milestoneID:        "invalid",
			requestBody:        UpdateMilestoneRequest{},
			mockSetup:          func(m *mockService.MockMilestoneService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando milestone no existe",
			milestoneID: "999",
			requestBody: UpdateMilestoneRequest{
				Title: &updatedTitle,
			},
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetMilestoneByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:        "retorna error cuando update falla",
			milestoneID: "1",
			requestBody: UpdateMilestoneRequest{
				Title: &updatedTitle,
			},
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					GetMilestoneByID(gomock.Any(), 1).
					Return(&milestone.Milestone{
						ID:        1,
						ProjectID: 1,
						Title:     &title,
						CreatedAt: now,
					}, nil)
				m.EXPECT().
					UpdateMilestone(gomock.Any(), gomock.Any()).
					Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockMilestoneService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/milestones/:id", controller.UpdateMilestone)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("PUT", "/milestones/"+tt.milestoneID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteMilestone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		milestoneID        string
		mockSetup          func(*mockService.MockMilestoneService)
		expectedStatusCode int
	}{
		{
			name:        "elimina milestone exitosamente",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					DeleteMilestone(gomock.Any(), 1).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			milestoneID:        "invalid",
			mockSetup:          func(m *mockService.MockMilestoneService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando delete falla",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockMilestoneService) {
				m.EXPECT().
					DeleteMilestone(gomock.Any(), 1).
					Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockMilestoneService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/milestones/:id", controller.DeleteMilestone)

			req, _ := http.NewRequest("DELETE", "/milestones/"+tt.milestoneID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

package feedback

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

	"softpharos/internal/core/domain/feedback"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllFeedbacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockFeedbackService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los feedbacks exitosamente",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetAllFeedbacks(gomock.Any()).Return([]feedback.Feedback{
					{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Feedback 1", CreatedAt: now},
					{ID: 2, MilestoneID: 1, ProfessorID: 1, Content: "Feedback 2", CreatedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetAllFeedbacks(gomock.Any()).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockFeedbackService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/feedbacks", controller.GetAllFeedbacks)
			req, _ := http.NewRequest("GET", "/feedbacks", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetFeedbackByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name               string
		feedbackID         string
		mockSetup          func(*mockService.MockFeedbackService)
		expectedStatusCode int
	}{
		{
			name:       "retorna feedback exitosamente",
			feedbackID: "1",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetFeedbackByID(gomock.Any(), 1).Return(&feedback.Feedback{
					ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Test feedback", CreatedAt: now,
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			feedbackID:         "invalid",
			mockSetup:          func(m *mockService.MockFeedbackService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:       "retorna error cuando feedback no existe",
			feedbackID: "999",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetFeedbackByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockFeedbackService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/feedbacks/:id", controller.GetFeedbackByID)
			req, _ := http.NewRequest("GET", "/feedbacks/"+tt.feedbackID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetFeedbacksByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name               string
		milestoneID        string
		mockSetup          func(*mockService.MockFeedbackService)
		expectedStatusCode int
	}{
		{
			name:        "retorna feedbacks por milestoneID exitosamente",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetFeedbacksByMilestoneID(gomock.Any(), 1).Return([]feedback.Feedback{
					{ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Feedback", CreatedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para milestoneID inválido",
			milestoneID:        "invalid",
			mockSetup:          func(m *mockService.MockFeedbackService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetFeedbacksByMilestoneID(gomock.Any(), 1).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockFeedbackService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/feedbacks/milestone/:milestoneId", controller.GetFeedbacksByMilestoneID)
			req, _ := http.NewRequest("GET", "/feedbacks/milestone/"+tt.milestoneID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockFeedbackService)
		expectedStatusCode int
	}{
		{
			name:        "crea feedback exitosamente",
			requestBody: CreateFeedbackRequest{MilestoneID: 1, ProfessorID: 1, Content: "New feedback"},
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().CreateFeedback(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockFeedbackService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			requestBody: CreateFeedbackRequest{MilestoneID: 1, ProfessorID: 1, Content: "New feedback"},
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().CreateFeedback(gomock.Any(), gomock.Any()).Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockFeedbackService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/feedbacks", controller.CreateFeedback)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("POST", "/feedbacks", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name               string
		feedbackID         string
		requestBody        interface{}
		mockSetup          func(*mockService.MockFeedbackService)
		expectedStatusCode int
	}{
		{
			name:        "actualiza feedback exitosamente",
			feedbackID:  "1",
			requestBody: UpdateFeedbackRequest{Content: "Updated feedback"},
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetFeedbackByID(gomock.Any(), 1).Return(&feedback.Feedback{
					ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Old feedback", CreatedAt: now,
				}, nil)
				m.EXPECT().UpdateFeedback(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			feedbackID:         "invalid",
			requestBody:        UpdateFeedbackRequest{},
			mockSetup:          func(m *mockService.MockFeedbackService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando feedback no existe",
			feedbackID:  "999",
			requestBody: UpdateFeedbackRequest{Content: "Updated feedback"},
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetFeedbackByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:        "retorna error cuando update falla",
			feedbackID:  "1",
			requestBody: UpdateFeedbackRequest{Content: "Updated feedback"},
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().GetFeedbackByID(gomock.Any(), 1).Return(&feedback.Feedback{
					ID: 1, MilestoneID: 1, ProfessorID: 1, Content: "Old feedback", CreatedAt: now,
				}, nil)
				m.EXPECT().UpdateFeedback(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockFeedbackService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/feedbacks/:id", controller.UpdateFeedback)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("PUT", "/feedbacks/"+tt.feedbackID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		feedbackID         string
		mockSetup          func(*mockService.MockFeedbackService)
		expectedStatusCode int
	}{
		{
			name:       "elimina feedback exitosamente",
			feedbackID: "1",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().DeleteFeedback(gomock.Any(), 1).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			feedbackID:         "invalid",
			mockSetup:          func(m *mockService.MockFeedbackService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:       "retorna error cuando delete falla",
			feedbackID: "1",
			mockSetup: func(m *mockService.MockFeedbackService) {
				m.EXPECT().DeleteFeedback(gomock.Any(), 1).Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockFeedbackService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/feedbacks/:id", controller.DeleteFeedback)
			req, _ := http.NewRequest("DELETE", "/feedbacks/"+tt.feedbackID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

package deliverable

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

	"softpharos/internal/core/domain/deliverable"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllDeliverables(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockDeliverableService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los entregables exitosamente",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetAllDeliverables(gomock.Any()).Return([]deliverable.Deliverable{
					{ID: 1, MilestoneID: 1, URL: "http://url1.com", CreatedAt: now},
					{ID: 2, MilestoneID: 1, URL: "http://url2.com", CreatedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetAllDeliverables(gomock.Any()).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockDeliverableService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/deliverables", controller.GetAllDeliverables)
			req, _ := http.NewRequest("GET", "/deliverables", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetDeliverableByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name               string
		deliverableID      string
		mockSetup          func(*mockService.MockDeliverableService)
		expectedStatusCode int
	}{
		{
			name:          "retorna entregable exitosamente",
			deliverableID: "1",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetDeliverableByID(gomock.Any(), 1).Return(&deliverable.Deliverable{
					ID: 1, MilestoneID: 1, URL: "http://url.com", CreatedAt: now,
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			deliverableID:      "invalid",
			mockSetup:          func(m *mockService.MockDeliverableService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:          "retorna error cuando entregable no existe",
			deliverableID: "999",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetDeliverableByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockDeliverableService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/deliverables/:id", controller.GetDeliverableByID)
			req, _ := http.NewRequest("GET", "/deliverables/"+tt.deliverableID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetDeliverablesByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name               string
		milestoneID        string
		mockSetup          func(*mockService.MockDeliverableService)
		expectedStatusCode int
	}{
		{
			name:        "retorna entregables por milestoneID exitosamente",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetDeliverablesByMilestoneID(gomock.Any(), 1).Return([]deliverable.Deliverable{
					{ID: 1, MilestoneID: 1, URL: "http://url.com", CreatedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para milestoneID inválido",
			milestoneID:        "invalid",
			mockSetup:          func(m *mockService.MockDeliverableService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetDeliverablesByMilestoneID(gomock.Any(), 1).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockDeliverableService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/deliverables/milestone/:milestoneId", controller.GetDeliverablesByMilestoneID)
			req, _ := http.NewRequest("GET", "/deliverables/milestone/"+tt.milestoneID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateDeliverable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockDeliverableService)
		expectedStatusCode int
	}{
		{
			name:        "crea entregable exitosamente",
			requestBody: CreateDeliverableRequest{MilestoneID: 1, URL: "http://url.com"},
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().CreateDeliverable(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockDeliverableService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			requestBody: CreateDeliverableRequest{MilestoneID: 1, URL: "http://url.com"},
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().CreateDeliverable(gomock.Any(), gomock.Any()).Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockDeliverableService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/deliverables", controller.CreateDeliverable)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("POST", "/deliverables", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateDeliverable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	updatedURL := "http://updated.com"
	tests := []struct {
		name               string
		deliverableID      string
		requestBody        interface{}
		mockSetup          func(*mockService.MockDeliverableService)
		expectedStatusCode int
	}{
		{
			name:          "actualiza entregable exitosamente",
			deliverableID: "1",
			requestBody:   UpdateDeliverableRequest{URL: &updatedURL},
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetDeliverableByID(gomock.Any(), 1).Return(&deliverable.Deliverable{
					ID: 1, MilestoneID: 1, URL: "http://old.com", CreatedAt: now,
				}, nil)
				m.EXPECT().UpdateDeliverable(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			deliverableID:      "invalid",
			requestBody:        UpdateDeliverableRequest{},
			mockSetup:          func(m *mockService.MockDeliverableService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:          "retorna error cuando entregable no existe",
			deliverableID: "999",
			requestBody:   UpdateDeliverableRequest{URL: &updatedURL},
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetDeliverableByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:          "retorna error cuando update falla",
			deliverableID: "1",
			requestBody:   UpdateDeliverableRequest{URL: &updatedURL},
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().GetDeliverableByID(gomock.Any(), 1).Return(&deliverable.Deliverable{
					ID: 1, MilestoneID: 1, URL: "http://old.com", CreatedAt: now,
				}, nil)
				m.EXPECT().UpdateDeliverable(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockDeliverableService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/deliverables/:id", controller.UpdateDeliverable)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("PUT", "/deliverables/"+tt.deliverableID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteDeliverable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		deliverableID      string
		mockSetup          func(*mockService.MockDeliverableService)
		expectedStatusCode int
	}{
		{
			name:          "elimina entregable exitosamente",
			deliverableID: "1",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().DeleteDeliverable(gomock.Any(), 1).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			deliverableID:      "invalid",
			mockSetup:          func(m *mockService.MockDeliverableService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:          "retorna error cuando delete falla",
			deliverableID: "1",
			mockSetup: func(m *mockService.MockDeliverableService) {
				m.EXPECT().DeleteDeliverable(gomock.Any(), 1).Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockDeliverableService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/deliverables/:id", controller.DeleteDeliverable)
			req, _ := http.NewRequest("DELETE", "/deliverables/"+tt.deliverableID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

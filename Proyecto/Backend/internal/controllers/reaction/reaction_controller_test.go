package reaction

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

	"softpharos/internal/core/domain/reaction"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllReactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rType1 := "like"
	rType2 := "love"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockReactionService)
		expectedStatusCode int
	}{
		{
			name: "retorna todas las reacciones exitosamente",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetAllReactions(gomock.Any()).Return([]reaction.Reaction{
					{ID: 1, MilestoneID: 1, UserID: 1, Type: &rType1, CreatedAt: now},
					{ID: 2, MilestoneID: 1, UserID: 2, Type: &rType2, CreatedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetAllReactions(gomock.Any()).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockReactionService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/reactions", controller.GetAllReactions)
			req, _ := http.NewRequest("GET", "/reactions", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetReactionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rType := "like"
	now := time.Now()

	tests := []struct {
		name               string
		reactionID         string
		mockSetup          func(*mockService.MockReactionService)
		expectedStatusCode int
	}{
		{
			name:       "retorna reacción exitosamente",
			reactionID: "1",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetReactionByID(gomock.Any(), 1).Return(&reaction.Reaction{
					ID: 1, MilestoneID: 1, UserID: 1, Type: &rType, CreatedAt: now,
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			reactionID:         "invalid",
			mockSetup:          func(m *mockService.MockReactionService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:       "retorna error cuando reacción no existe",
			reactionID: "999",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetReactionByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockReactionService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/reactions/:id", controller.GetReactionByID)
			req, _ := http.NewRequest("GET", "/reactions/"+tt.reactionID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetReactionsByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rType := "like"
	now := time.Now()

	tests := []struct {
		name               string
		milestoneID        string
		mockSetup          func(*mockService.MockReactionService)
		expectedStatusCode int
	}{
		{
			name:        "retorna reacciones por milestoneID exitosamente",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetReactionsByMilestoneID(gomock.Any(), 1).Return([]reaction.Reaction{
					{ID: 1, MilestoneID: 1, UserID: 1, Type: &rType, CreatedAt: now},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para milestoneID inválido",
			milestoneID:        "invalid",
			mockSetup:          func(m *mockService.MockReactionService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetReactionsByMilestoneID(gomock.Any(), 1).Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockReactionService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/reactions/milestone/:milestoneId", controller.GetReactionsByMilestoneID)
			req, _ := http.NewRequest("GET", "/reactions/milestone/"+tt.milestoneID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rType := "like"

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockReactionService)
		expectedStatusCode int
	}{
		{
			name:        "crea reacción exitosamente",
			requestBody: CreateReactionRequest{MilestoneID: 1, UserID: 1, Type: &rType},
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().CreateReaction(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockReactionService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			requestBody: CreateReactionRequest{MilestoneID: 1, UserID: 1, Type: &rType},
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().CreateReaction(gomock.Any(), gomock.Any()).Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockReactionService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/reactions", controller.CreateReaction)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("POST", "/reactions", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rType := "like"
	updatedType := "love"
	now := time.Now()

	tests := []struct {
		name               string
		reactionID         string
		requestBody        interface{}
		mockSetup          func(*mockService.MockReactionService)
		expectedStatusCode int
	}{
		{
			name:        "actualiza reacción exitosamente",
			reactionID:  "1",
			requestBody: UpdateReactionRequest{Type: &updatedType},
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetReactionByID(gomock.Any(), 1).Return(&reaction.Reaction{
					ID: 1, MilestoneID: 1, UserID: 1, Type: &rType, CreatedAt: now,
				}, nil)
				m.EXPECT().UpdateReaction(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			reactionID:         "invalid",
			requestBody:        UpdateReactionRequest{},
			mockSetup:          func(m *mockService.MockReactionService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando reacción no existe",
			reactionID:  "999",
			requestBody: UpdateReactionRequest{Type: &updatedType},
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetReactionByID(gomock.Any(), 999).Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:        "retorna error cuando update falla",
			reactionID:  "1",
			requestBody: UpdateReactionRequest{Type: &updatedType},
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().GetReactionByID(gomock.Any(), 1).Return(&reaction.Reaction{
					ID: 1, MilestoneID: 1, UserID: 1, Type: &rType, CreatedAt: now,
				}, nil)
				m.EXPECT().UpdateReaction(gomock.Any(), gomock.Any()).Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockReactionService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/reactions/:id", controller.UpdateReaction)
			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req, _ := http.NewRequest("PUT", "/reactions/"+tt.reactionID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		reactionID         string
		mockSetup          func(*mockService.MockReactionService)
		expectedStatusCode int
	}{
		{
			name:       "elimina reacción exitosamente",
			reactionID: "1",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().DeleteReaction(gomock.Any(), 1).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			reactionID:         "invalid",
			mockSetup:          func(m *mockService.MockReactionService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:       "retorna error cuando delete falla",
			reactionID: "1",
			mockSetup: func(m *mockService.MockReactionService) {
				m.EXPECT().DeleteReaction(gomock.Any(), 1).Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockReactionService(ctrl)
			tt.mockSetup(mockSvc)
			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/reactions/:id", controller.DeleteReaction)
			req, _ := http.NewRequest("DELETE", "/reactions/"+tt.reactionID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

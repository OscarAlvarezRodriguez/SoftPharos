package comment

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

	"softpharos/internal/core/domain/comment"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content1 := "Comment 1"
	content2 := "Comment 2"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockCommentService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los comentarios exitosamente",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetAllComments(gomock.Any()).
					Return([]comment.Comment{
						{ID: 1, MilestoneID: 1, UserID: 1, Content: &content1, CreatedAt: now},
						{ID: 2, MilestoneID: 1, UserID: 2, Content: &content2, CreatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetAllComments(gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockCommentService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/comments", controller.GetAllComments)

			req, _ := http.NewRequest("GET", "/comments", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetCommentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content := "Test Comment"
	now := time.Now()

	tests := []struct {
		name               string
		commentID          string
		mockSetup          func(*mockService.MockCommentService)
		expectedStatusCode int
	}{
		{
			name:      "retorna comentario exitosamente",
			commentID: "1",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetCommentByID(gomock.Any(), 1).
					Return(&comment.Comment{
						ID:          1,
						MilestoneID: 1,
						UserID:      1,
						Content:     &content,
						CreatedAt:   now,
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			commentID:          "invalid",
			mockSetup:          func(m *mockService.MockCommentService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando comentario no existe",
			commentID: "999",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetCommentByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockCommentService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/comments/:id", controller.GetCommentByID)

			req, _ := http.NewRequest("GET", "/comments/"+tt.commentID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetCommentsByMilestoneID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content1 := "Comment 1"
	content2 := "Comment 2"
	now := time.Now()

	tests := []struct {
		name               string
		milestoneID        string
		mockSetup          func(*mockService.MockCommentService)
		expectedStatusCode int
	}{
		{
			name:        "retorna comentarios por milestoneID exitosamente",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetCommentsByMilestoneID(gomock.Any(), 1).
					Return([]comment.Comment{
						{ID: 1, MilestoneID: 1, UserID: 1, Content: &content1, CreatedAt: now},
						{ID: 2, MilestoneID: 1, UserID: 2, Content: &content2, CreatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para milestoneID inválido",
			milestoneID:        "invalid",
			mockSetup:          func(m *mockService.MockCommentService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "retorna error cuando el service falla",
			milestoneID: "1",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetCommentsByMilestoneID(gomock.Any(), 1).
					Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockCommentService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/comments/milestone/:milestoneId", controller.GetCommentsByMilestoneID)

			req, _ := http.NewRequest("GET", "/comments/milestone/"+tt.milestoneID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content := "New Comment"

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockCommentService)
		expectedStatusCode int
	}{
		{
			name: "crea comentario exitosamente",
			requestBody: CreateCommentRequest{
				MilestoneID: 1,
				UserID:      1,
				Content:     &content,
			},
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					CreateComment(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockCommentService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "retorna error cuando el service falla",
			requestBody: CreateCommentRequest{
				MilestoneID: 1,
				UserID:      1,
				Content:     &content,
			},
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					CreateComment(gomock.Any(), gomock.Any()).
					Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockCommentService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/comments", controller.CreateComment)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("POST", "/comments", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	content := "Existing Comment"
	updatedContent := "Updated Comment"
	now := time.Now()

	tests := []struct {
		name               string
		commentID          string
		requestBody        interface{}
		mockSetup          func(*mockService.MockCommentService)
		expectedStatusCode int
	}{
		{
			name:      "actualiza comentario exitosamente",
			commentID: "1",
			requestBody: UpdateCommentRequest{
				Content: &updatedContent,
			},
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetCommentByID(gomock.Any(), 1).
					Return(&comment.Comment{
						ID:          1,
						MilestoneID: 1,
						UserID:      1,
						Content:     &content,
						CreatedAt:   now,
					}, nil)
				m.EXPECT().
					UpdateComment(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			commentID:          "invalid",
			requestBody:        UpdateCommentRequest{},
			mockSetup:          func(m *mockService.MockCommentService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando comentario no existe",
			commentID: "999",
			requestBody: UpdateCommentRequest{
				Content: &updatedContent,
			},
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetCommentByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:      "retorna error cuando update falla",
			commentID: "1",
			requestBody: UpdateCommentRequest{
				Content: &updatedContent,
			},
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					GetCommentByID(gomock.Any(), 1).
					Return(&comment.Comment{
						ID:          1,
						MilestoneID: 1,
						UserID:      1,
						Content:     &content,
						CreatedAt:   now,
					}, nil)
				m.EXPECT().
					UpdateComment(gomock.Any(), gomock.Any()).
					Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockCommentService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/comments/:id", controller.UpdateComment)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("PUT", "/comments/"+tt.commentID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		commentID          string
		mockSetup          func(*mockService.MockCommentService)
		expectedStatusCode int
	}{
		{
			name:      "elimina comentario exitosamente",
			commentID: "1",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					DeleteComment(gomock.Any(), 1).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			commentID:          "invalid",
			mockSetup:          func(m *mockService.MockCommentService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "retorna error cuando delete falla",
			commentID: "1",
			mockSetup: func(m *mockService.MockCommentService) {
				m.EXPECT().
					DeleteComment(gomock.Any(), 1).
					Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockCommentService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/comments/:id", controller.DeleteComment)

			req, _ := http.NewRequest("DELETE", "/comments/"+tt.commentID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

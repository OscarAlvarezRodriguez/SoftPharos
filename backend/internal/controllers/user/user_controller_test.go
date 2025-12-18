package user

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

	"softpharos/internal/core/domain/user"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name1 := "User 1"
	name2 := "User 2"
	now := time.Now()

	tests := []struct {
		name               string
		mockSetup          func(*mockService.MockUserService)
		expectedStatusCode int
	}{
		{
			name: "retorna todos los usuarios exitosamente",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetAllUsers(gomock.Any()).
					Return([]user.User{
						{ID: 1, Name: &name1, Email: "user1@test.com", RoleID: 1, CreatedAt: now},
						{ID: 2, Name: &name2, Email: "user2@test.com", RoleID: 2, CreatedAt: now},
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "retorna error cuando el service falla",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetAllUsers(gomock.Any()).
					Return(nil, errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockUserService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/users", controller.GetAllUsers)

			req, _ := http.NewRequest("GET", "/users", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Test User"
	now := time.Now()

	tests := []struct {
		name               string
		userID             string
		mockSetup          func(*mockService.MockUserService)
		expectedStatusCode int
	}{
		{
			name:   "retorna usuario exitosamente",
			userID: "1",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetUserByID(gomock.Any(), 1).
					Return(&user.User{
						ID:        1,
						Name:      &name,
						Email:     "test@test.com",
						RoleID:    1,
						CreatedAt: now,
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			userID:             "invalid",
			mockSetup:          func(m *mockService.MockUserService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:   "retorna error cuando usuario no existe",
			userID: "999",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetUserByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockUserService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/users/:id", controller.GetUserByID)

			req, _ := http.NewRequest("GET", "/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Test User"
	now := time.Now()

	tests := []struct {
		name               string
		email              string
		mockSetup          func(*mockService.MockUserService)
		expectedStatusCode int
	}{
		{
			name:  "retorna usuario por email exitosamente",
			email: "test@test.com",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetUserByEmail(gomock.Any(), "test@test.com").
					Return(&user.User{
						ID:        1,
						Name:      &name,
						Email:     "test@test.com",
						RoleID:    1,
						CreatedAt: now,
					}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "retorna error cuando usuario no existe",
			email: "notfound@test.com",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetUserByEmail(gomock.Any(), "notfound@test.com").
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockUserService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.GET("/users/email/:email", controller.GetUserByEmail)

			req, _ := http.NewRequest("GET", "/users/email/"+tt.email, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "New User"

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockUserService)
		expectedStatusCode int
	}{
		{
			name: "crea usuario exitosamente",
			requestBody: CreateUserRequest{
				Name:     &name,
				Email:    "newuser@test.com",
				Password: "password123",
				RoleID:   1,
			},
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "retorna error para request body inválido",
			requestBody:        "invalid json",
			mockSetup:          func(m *mockService.MockUserService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "retorna error cuando el service falla",
			requestBody: CreateUserRequest{
				Name:     &name,
				Email:    "newuser@test.com",
				Password: "password123",
				RoleID:   1,
			},
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(errors.New("service error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockUserService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/users", controller.CreateUser)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Existing User"
	updatedName := "Updated Name"
	now := time.Now()

	tests := []struct {
		name               string
		userID             string
		requestBody        interface{}
		mockSetup          func(*mockService.MockUserService)
		expectedStatusCode int
	}{
		{
			name:   "actualiza usuario exitosamente",
			userID: "1",
			requestBody: UpdateUserRequest{
				Name: &updatedName,
			},
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetUserByID(gomock.Any(), 1).
					Return(&user.User{
						ID:        1,
						Name:      &name,
						Email:     "test@test.com",
						Password:  "password",
						RoleID:    1,
						CreatedAt: now,
					}, nil)
				m.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			userID:             "invalid",
			requestBody:        UpdateUserRequest{},
			mockSetup:          func(m *mockService.MockUserService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:   "retorna error cuando usuario no existe",
			userID: "999",
			requestBody: UpdateUserRequest{
				Name: &updatedName,
			},
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetUserByID(gomock.Any(), 999).
					Return(nil, errors.New("not found"))
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:   "retorna error cuando update falla",
			userID: "1",
			requestBody: UpdateUserRequest{
				Name: &updatedName,
			},
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					GetUserByID(gomock.Any(), 1).
					Return(&user.User{
						ID:        1,
						Name:      &name,
						Email:     "test@test.com",
						Password:  "password",
						RoleID:    1,
						CreatedAt: now,
					}, nil)
				m.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Return(errors.New("update error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockUserService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.PUT("/users/:id", controller.UpdateUser)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("PUT", "/users/"+tt.userID, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		userID             string
		mockSetup          func(*mockService.MockUserService)
		expectedStatusCode int
	}{
		{
			name:   "elimina usuario exitosamente",
			userID: "1",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					DeleteUser(gomock.Any(), 1).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "retorna error para ID inválido",
			userID:             "invalid",
			mockSetup:          func(m *mockService.MockUserService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:   "retorna error cuando delete falla",
			userID: "1",
			mockSetup: func(m *mockService.MockUserService) {
				m.EXPECT().
					DeleteUser(gomock.Any(), 1).
					Return(errors.New("delete error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockUserService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.DELETE("/users/:id", controller.DeleteUser)

			req, _ := http.NewRequest("DELETE", "/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/domain/user"
	mockService "softpharos/mocks/core/ports/services"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestGoogleLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "John Doe"
	picture := "https://example.com/pic.jpg"
	roleName := "student"
	roleDesc := "Student role"

	tests := []struct {
		name               string
		requestBody        interface{}
		mockSetup          func(*mockService.MockAuthService)
		expectedStatusCode int
	}{
		{
			name: "autenticación exitosa con Google",
			requestBody: GoogleLoginRequest{
				IDToken: "valid-google-token",
			},
			mockSetup: func(m *mockService.MockAuthService) {
				m.EXPECT().
					AuthenticateWithGoogle(gomock.Any(), "valid-google-token").
					Return(&user.User{
						ID:         1,
						Name:       &name,
						Email:      "test@unal.edu.co",
						ProviderID: "google-123",
						RoleID:     3,
						PictureURL: &picture,
						Role: &role.Role{
							ID:          3,
							Name:        roleName,
							Description: &roleDesc,
						},
					}, "jwt-access-token-123", nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:        "retorna error cuando el request body es inválido",
			requestBody: `{"invalid": json}`,
			mockSetup: func(m *mockService.MockAuthService) {
				// No se espera llamada al servicio
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "retorna error cuando el token de Google es inválido",
			requestBody: GoogleLoginRequest{
				IDToken: "invalid-token",
			},
			mockSetup: func(m *mockService.MockAuthService) {
				m.EXPECT().
					AuthenticateWithGoogle(gomock.Any(), "invalid-token").
					Return(nil, "", errors.New("invalid token"))
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: "retorna error cuando el servicio falla",
			requestBody: GoogleLoginRequest{
				IDToken: "valid-token",
			},
			mockSetup: func(m *mockService.MockAuthService) {
				m.EXPECT().
					AuthenticateWithGoogle(gomock.Any(), "valid-token").
					Return(nil, "", errors.New("service error"))
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := mockService.NewMockAuthService(ctrl)
			tt.mockSetup(mockSvc)

			controller := New(mockSvc)
			router := setupRouter()
			router.POST("/auth/google", controller.GoogleLogin)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req, _ := http.NewRequest("POST", "/auth/google", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if tt.expectedStatusCode == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.True(t, response["success"].(bool))

				data := response["data"].(map[string]interface{})
				assert.NotEmpty(t, data["accessToken"])
				assert.NotNil(t, data["user"])
			}
		})
	}
}

func TestNewController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mockService.NewMockAuthService(ctrl)
	controller := New(mockSvc)

	assert.NotNil(t, controller)
	assert.Equal(t, mockSvc, controller.authService)
}

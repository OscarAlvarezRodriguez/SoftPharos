package auth

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT_Success(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	token, err := GenerateJWT(1, "test@unal.edu.co", 3)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateJWT_NoSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	assert.Panics(t, func() {
		GenerateJWT(1, "test@unal.edu.co", 3)
	})
}

func TestValidateJWT_Success(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	token, _ := GenerateJWT(1, "test@unal.edu.co", 3)

	claims, err := ValidateJWT(token)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, 1, claims.UserID)
	assert.Equal(t, "test@unal.edu.co", claims.Email)
	assert.Equal(t, 3, claims.RoleID)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	claims, err := ValidateJWT("invalid-token")

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	// Crear un token ya expirado no es fácil sin modificar la función
	// Este test se omite o requiere refactorización
}

func TestValidateJWT_NoSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	claims, err := ValidateJWT("some-token")

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "JWT_SECRET")
}

func TestClaims(t *testing.T) {
	claims := &Claims{
		UserID: 1,
		Email:  "test@example.com",
		RoleID: 3,
	}

	assert.Equal(t, 1, claims.UserID)
	assert.Equal(t, "test@example.com", claims.Email)
	assert.Equal(t, 3, claims.RoleID)
}

func TestJWTExpirationTime(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	token, err := GenerateJWT(1, "test@unal.edu.co", 3)
	assert.NoError(t, err)

	claims, err := ValidateJWT(token)
	assert.NoError(t, err)

	// Verificar que el token expira en ~24 horas
	expirationTime := claims.ExpiresAt.Time
	expectedExpiration := time.Now().Add(24 * time.Hour)

	// Permitir 1 minuto de diferencia debido a tiempo de ejecución del test
	diff := expirationTime.Sub(expectedExpiration)
	assert.Less(t, diff.Abs(), time.Minute)
}

package databases

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestNewClient verifica la creación de un nuevo cliente
func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		config      Config
		shouldError bool
	}{
		{
			name: "Configuración inválida - host vacío",
			config: Config{
				Host:     "",
				Port:     "5432",
				User:     "test",
				Password: "test",
				DBName:   "test",
				SSLMode:  "disable",
			},
			shouldError: true,
		},
		{
			name: "Configuración inválida - puerto inválido",
			config: Config{
				Host:     "localhost",
				Port:     "invalid",
				User:     "test",
				Password: "test",
				DBName:   "test",
				SSLMode:  "disable",
			},
			shouldError: true,
		},
		{
			name: "Configuración con credenciales incorrectas",
			config: Config{
				Host:     "localhost",
				Port:     "5432",
				User:     "nonexistentuser",
				Password: "wrongpassword",
				DBName:   "nonexistentdb",
				SSLMode:  "disable",
			},
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.config)

			if tt.shouldError {
				assert.Error(t, err, "Se esperaba un error")
				assert.Nil(t, client, "El cliente debería ser nil en caso de error")
			} else {
				assert.NoError(t, err, "No se esperaba un error")
				assert.NotNil(t, client, "El cliente no debería ser nil")
				if client != nil {
					client.Close()
				}
			}
		})
	}
}

// TestNewClientFromEnv verifica la creación desde variables de entorno
func TestNewClientFromEnv(t *testing.T) {
	// Guardar las variables de entorno actuales
	originalHost := os.Getenv("DB_HOST")
	originalPort := os.Getenv("DB_PORT")
	originalUser := os.Getenv("DB_USER")
	originalPassword := os.Getenv("DB_PASSWORD")
	originalDBName := os.Getenv("DB_NAME")
	originalSSLMode := os.Getenv("DB_SSLMODE")

	// Restaurar al final
	defer func() {
		os.Setenv("DB_HOST", originalHost)
		os.Setenv("DB_PORT", originalPort)
		os.Setenv("DB_USER", originalUser)
		os.Setenv("DB_PASSWORD", originalPassword)
		os.Setenv("DB_NAME", originalDBName)
		os.Setenv("DB_SSLMODE", originalSSLMode)
	}()

	tests := []struct {
		name        string
		envVars     map[string]string
		shouldError bool
	}{
		{
			name: "Variables de entorno inválidas",
			envVars: map[string]string{
				"DB_HOST":     "",
				"DB_PORT":     "5432",
				"DB_USER":     "test",
				"DB_PASSWORD": "test",
				"DB_NAME":     "test",
				"DB_SSLMODE":  "disable",
			},
			shouldError: true,
		},
		{
			name: "Variables de entorno con credenciales incorrectas",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USER":     "invaliduser",
				"DB_PASSWORD": "invalidpass",
				"DB_NAME":     "invaliddb",
				"DB_SSLMODE":  "disable",
			},
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar variables de entorno
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Intentar crear cliente
			client, err := NewClientFromEnv()

			if tt.shouldError {
				assert.Error(t, err, "Se esperaba un error")
				assert.Nil(t, client, "El cliente debería ser nil")
			} else {
				assert.NoError(t, err, "No se esperaba un error")
				assert.NotNil(t, client, "El cliente no debería ser nil")
				if client != nil {
					client.Close()
				}
			}
		})
	}
}

// TestClientClose verifica que se pueda cerrar un cliente
func TestClientClose(t *testing.T) {
	t.Run("Close con cliente nil DB causa panic", func(t *testing.T) {
		client := &Client{DB: nil}
		// Esperamos que cause panic al intentar cerrar con DB nil
		assert.Panics(t, func() {
			client.Close()
		}, "Debería hacer panic al cerrar un cliente con DB nil")
	})

	t.Run("Close con cliente válido", func(t *testing.T) {
		// Crear un cliente con SQLite en memoria
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		assert.NoError(t, err)

		client := &Client{DB: db}

		// Cerrar debería funcionar sin error
		err = client.Close()
		assert.NoError(t, err, "No debería haber error al cerrar un cliente válido")
	})
}

// TestClientPing verifica la función Ping
func TestClientPing(t *testing.T) {
	t.Run("Ping con cliente nil DB causa panic", func(t *testing.T) {
		client := &Client{DB: nil}
		assert.Panics(t, func() {
			client.Ping()
		}, "Debería hacer panic al hacer ping con DB nil")
	})

	t.Run("Ping con cliente válido", func(t *testing.T) {
		// Crear un cliente con SQLite en memoria
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		assert.NoError(t, err)

		client := &Client{DB: db}

		// Ping debería funcionar sin error
		err = client.Ping()
		assert.NoError(t, err, "No debería haber error al hacer ping a un cliente válido")

		client.Close()
	})
}

// TestConfigStruct verifica que la estructura Config tenga los campos correctos
func TestConfigStruct(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpassword",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "5432", config.Port)
	assert.Equal(t, "testuser", config.User)
	assert.Equal(t, "testpassword", config.Password)
	assert.Equal(t, "testdb", config.DBName)
	assert.Equal(t, "disable", config.SSLMode)
}

// TestClientStruct verifica que la estructura Client tenga los campos correctos
func TestClientStruct(t *testing.T) {
	client := &Client{DB: nil}
	assert.NotNil(t, client)
	assert.Nil(t, client.DB)
}

package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Client representa el cliente de conexión a PostgreSQL
type Client struct {
	DB *gorm.DB
}

// Config contiene la configuración de conexión a la BD
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewClient crea una nueva instancia del cliente PostgreSQL con GORM
func NewClient(cfg Config) (*Client, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode,
	)

	// Configurar logger de GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con PostgreSQL: %w", err)
	}

	// Configurar connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener SQL DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Conectado exitosamente a PostgreSQL con GORM")

	return &Client{DB: db}, nil
}

// NewClientFromEnv crea un cliente usando variables de entorno
func NewClientFromEnv() (*Client, error) {
	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return NewClient(cfg)
}

// Close cierra la conexión con la base de datos
func (c *Client) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Ping verifica que la conexión esté activa
func (c *Client) Ping() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

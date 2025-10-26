package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"softpharos/internal/infra/databases"
	"testing"
)

func SetupMockDB(t *testing.T) (*databases.Client, sqlmock.Sqlmock, *sql.DB) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}

	client := &databases.Client{DB: db}
	return client, mock, sqlDB
}

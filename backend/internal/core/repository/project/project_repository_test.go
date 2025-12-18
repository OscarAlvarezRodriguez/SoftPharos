package project

import (
	"context"
	"errors"
	"regexp"
	"softpharos/internal/core/domain/project"
	"softpharos/internal/core/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetAll(t *testing.T) {
	name1 := "Project 1"
	name2 := "Project 2"
	ownerName1 := "Owner 1"
	ownerName2 := "Owner 2"
	now := time.Now()

	tests := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectedLen   int
		expectedError bool
	}{
		{
			name: "retorna todos los proyectos exitosamente",
			mockSetup: func(mock sqlmock.Sqlmock) {
				projectRows := sqlmock.NewRows([]string{"id", "name", "objective", "created_by", "created_at", "updated_at"}).
					AddRow(1, name1, nil, 1, now, now).
					AddRow(2, name2, nil, 2, now, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project"`)).
					WillReturnRows(projectRows)

				ownerRows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at"}).
					AddRow(1, ownerName1, "owner1@example.com", "hash1", 1, now).
					AddRow(2, ownerName2, "owner2@example.com", "hash2", 1, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "user"."id" IN ($1,$2)`)).
					WithArgs(1, 2).
					WillReturnRows(ownerRows)
			},
			expectedLen:   2,
			expectedError: false,
		},
		{
			name: "retorna error cuando la query falla",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project"`)).
					WillReturnError(errors.New("database error"))
			},
			expectedLen:   0,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mock, sqlDB := repository.SetupMockDB(t)
			defer sqlDB.Close()

			tt.mockSetup(mock)

			repo := New(client)
			ctx := context.Background()

			projects, err := repo.GetAll(ctx)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, projects, tt.expectedLen)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetByID(t *testing.T) {
	name := "Test Project"
	ownerName := "Owner Name"
	now := time.Now()

	tests := []struct {
		name          string
		projectID     int
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:      "retorna proyecto por ID exitosamente",
			projectID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				projectRows := sqlmock.NewRows([]string{"id", "name", "objective", "created_by", "created_at", "updated_at"}).
					AddRow(1, name, nil, 1, now, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project" WHERE "project"."id" = $1`)).
					WithArgs(1, 1).
					WillReturnRows(projectRows)

				ownerRows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at"}).
					AddRow(1, ownerName, "owner@example.com", "hash", 1, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "user"."id" = $1`)).
					WithArgs(1).
					WillReturnRows(ownerRows)
			},
			expectedError: false,
		},
		{
			name:      "retorna error cuando el proyecto no existe",
			projectID: 999,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project" WHERE "project"."id" = $1`)).
					WithArgs(999, 1).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mock, sqlDB := repository.SetupMockDB(t)
			defer sqlDB.Close()

			tt.mockSetup(mock)

			repo := New(client)
			ctx := context.Background()

			project, err := repo.GetByID(ctx, tt.projectID)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, project)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, project)
				assert.Equal(t, tt.projectID, project.ID)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCreate(t *testing.T) {
	name := "New Project"

	tests := []struct {
		name          string
		project       *project.Project
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:    "crea proyecto exitosamente",
			project: &project.Project{Name: &name, CreatedBy: 1},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "project"`)).
					WithArgs(name, nil, 1, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
						AddRow(1, time.Now(), time.Now()))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name:    "retorna error cuando create falla",
			project: &project.Project{Name: &name, CreatedBy: 1},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "project"`)).
					WithArgs(name, nil, 1, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mock, sqlDB := repository.SetupMockDB(t)
			defer sqlDB.Close()

			tt.mockSetup(mock)

			repo := New(client)
			ctx := context.Background()

			err := repo.Create(ctx, tt.project)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.project.ID)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdate(t *testing.T) {
	name := "Updated Project"

	tests := []struct {
		name          string
		project       *project.Project
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name: "actualiza proyecto exitosamente",
			project: &project.Project{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "project" SET`)).
					WithArgs(name, nil, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name: "retorna error cuando update falla",
			project: &project.Project{
				ID:        1,
				Name:      &name,
				CreatedBy: 1,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "project" SET`)).
					WithArgs(name, nil, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mock, sqlDB := repository.SetupMockDB(t)
			defer sqlDB.Close()

			tt.mockSetup(mock)

			repo := New(client)
			ctx := context.Background()

			err := repo.Update(ctx, tt.project)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name          string
		projectID     int
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:      "elimina proyecto exitosamente",
			projectID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "project" WHERE "project"."id" = $1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name:      "retorna error cuando delete falla",
			projectID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "project" WHERE "project"."id" = $1`)).
					WithArgs(1).
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mock, sqlDB := repository.SetupMockDB(t)
			defer sqlDB.Close()

			tt.mockSetup(mock)

			repo := New(client)
			ctx := context.Background()

			err := repo.Delete(ctx, tt.projectID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

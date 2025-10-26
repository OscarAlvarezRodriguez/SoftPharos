package role

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/repository"
)

func TestRoleGetAll(t *testing.T) {
	desc1 := "Admin role"
	desc2 := "User role"
	now := time.Now()

	tests := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectedLen   int
		expectedError bool
	}{
		{
			name: "retorna todos los roles exitosamente",
			mockSetup: func(mock sqlmock.Sqlmock) {
				roleRows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
					AddRow(1, "Admin", desc1, now).
					AddRow(2, "User", desc2, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role"`)).
					WillReturnRows(roleRows)
			},
			expectedLen:   2,
			expectedError: false,
		},
		{
			name: "retorna error cuando la query falla",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role"`)).
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

			roles, err := repo.GetAll(ctx)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, roles, tt.expectedLen)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRoleGetByID(t *testing.T) {
	desc := "Admin role"
	now := time.Now()

	tests := []struct {
		name          string
		roleID        int
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:   "retorna role por ID exitosamente",
			roleID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				roleRows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
					AddRow(1, "Admin", desc, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role" WHERE "role"."id" = $1`)).
					WithArgs(1, 1).
					WillReturnRows(roleRows)
			},
			expectedError: false,
		},
		{
			name:   "retorna error cuando el role no existe",
			roleID: 999,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role" WHERE "role"."id" = $1`)).
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

			role, err := repo.GetByID(ctx, tt.roleID)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, role)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, role)
				assert.Equal(t, tt.roleID, role.ID)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRoleGetByName(t *testing.T) {
	desc := "Admin role"
	now := time.Now()

	tests := []struct {
		name          string
		roleName      string
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:     "retorna role por nombre exitosamente",
			roleName: "Admin",
			mockSetup: func(mock sqlmock.Sqlmock) {
				roleRows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
					AddRow(1, "Admin", desc, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role" WHERE name = $1`)).
					WithArgs("Admin", 1).
					WillReturnRows(roleRows)
			},
			expectedError: false,
		},
		{
			name:     "retorna error cuando el role no existe",
			roleName: "NonExistent",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role" WHERE name = $1`)).
					WithArgs("NonExistent", 1).
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

			role, err := repo.GetByName(ctx, tt.roleName)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, role)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, role)
				assert.Equal(t, tt.roleName, role.Name)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRoleCreate(t *testing.T) {
	desc := "New role description"

	tests := []struct {
		name          string
		role          *role.Role
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name: "crea role exitosamente",
			role: &role.Role{Name: "NewRole", Description: &desc},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "role"`)).
					WithArgs("NewRole", desc, sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
						AddRow(1, time.Now()))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name: "retorna error cuando create falla",
			role: &role.Role{Name: "NewRole", Description: &desc},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "role"`)).
					WithArgs("NewRole", desc, sqlmock.AnyArg()).
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

			err := repo.Create(ctx, tt.role)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.role.ID)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRoleUpdate(t *testing.T) {
	desc := "Updated role description"

	tests := []struct {
		name          string
		role          *role.Role
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name: "actualiza role exitosamente",
			role: &role.Role{
				ID:          1,
				Name:        "UpdatedRole",
				Description: &desc,
				CreatedAt:   time.Now(),
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "role" SET`)).
					WithArgs("UpdatedRole", desc, sqlmock.AnyArg(), 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name: "retorna error cuando update falla",
			role: &role.Role{
				ID:          1,
				Name:        "UpdatedRole",
				Description: &desc,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "role" SET`)).
					WithArgs("UpdatedRole", desc, sqlmock.AnyArg(), 1).
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

			err := repo.Update(ctx, tt.role)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRoleDelete(t *testing.T) {
	tests := []struct {
		name          string
		roleID        int
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:   "elimina role exitosamente",
			roleID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "role" WHERE "role"."id" = $1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name:   "retorna error cuando delete falla",
			roleID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "role" WHERE "role"."id" = $1`)).
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

			err := repo.Delete(ctx, tt.roleID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

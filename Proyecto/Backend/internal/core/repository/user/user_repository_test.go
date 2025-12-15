package user

import (
	"context"
	"errors"
	"regexp"
	"softpharos/internal/core/domain/user"
	"softpharos/internal/core/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetAll(t *testing.T) {
	name1 := "User 1"
	name2 := "User 2"
	roleName := "Admin"
	now := time.Now()

	tests := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectedLen   int
		expectedError bool
	}{
		{
			name: "retorna todos los usuarios exitosamente",
			mockSetup: func(mock sqlmock.Sqlmock) {
				userRows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at"}).
					AddRow(1, name1, "user1@example.com", "hash1", 1, now).
					AddRow(2, name2, "user2@example.com", "hash2", 2, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user"`)).
					WillReturnRows(userRows)

				roleRows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
					AddRow(1, roleName, nil, now).
					AddRow(2, roleName, nil, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role" WHERE "role"."id" IN ($1,$2)`)).
					WithArgs(1, 2).
					WillReturnRows(roleRows)
			},
			expectedLen:   2,
			expectedError: false,
		},
		{
			name: "retorna error cuando la query falla",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user"`)).
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

			users, err := repo.GetAll(ctx)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, users, tt.expectedLen)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetByID(t *testing.T) {
	name := "Test User"
	roleName := "Admin"
	now := time.Now()

	tests := []struct {
		name          string
		userID        int
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:   "retorna usuario por ID exitosamente",
			userID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				userRows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at"}).
					AddRow(1, name, "test@example.com", "hash", 1, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "user"."id" = $1`)).
					WithArgs(1, 1).
					WillReturnRows(userRows)

				roleRows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
					AddRow(1, roleName, nil, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role" WHERE "role"."id" = $1`)).
					WithArgs(1).
					WillReturnRows(roleRows)
			},
			expectedError: false,
		},
		{
			name:   "retorna error cuando el usuario no existe",
			userID: 999,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "user"."id" = $1`)).
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

			user, err := repo.GetByID(ctx, tt.userID)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.userID, user.ID)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetByEmail(t *testing.T) {
	name := "Test User"
	email := "test@example.com"
	roleName := "Admin"
	now := time.Now()

	tests := []struct {
		name          string
		email         string
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:  "retorna usuario por email exitosamente",
			email: email,
			mockSetup: func(mock sqlmock.Sqlmock) {
				userRows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_id", "created_at"}).
					AddRow(1, name, email, "hash", 1, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE email = $1`)).
					WithArgs(email, 1).
					WillReturnRows(userRows)

				roleRows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
					AddRow(1, roleName, nil, now)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role" WHERE "role"."id" = $1`)).
					WithArgs(1).
					WillReturnRows(roleRows)
			},
			expectedError: false,
		},
		{
			name:  "retorna error cuando el email no existe",
			email: "notfound@example.com",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE email = $1`)).
					WithArgs("notfound@example.com", 1).
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

			user, err := repo.GetByEmail(ctx, tt.email)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.email, user.Email)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCreate(t *testing.T) {
	name := "New User"

	tests := []struct {
		name          string
		user          *user.User
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name: "crea usuario exitosamente",
			user: &user.User{Name: &name, Email: "new@example.com", Password: "hash", RoleID: 1},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "user"`)).
					WithArgs(name, "new@example.com", "hash", 1, sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
						AddRow(1, time.Now()))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name: "retorna error cuando create falla",
			user: &user.User{Name: &name, Email: "new@example.com", Password: "hash", RoleID: 1},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "user"`)).
					WithArgs(name, "new@example.com", "hash", 1, sqlmock.AnyArg()).
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

			err := repo.Create(ctx, tt.user)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.user.ID)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdate(t *testing.T) {
	name := "Updated User"

	tests := []struct {
		name          string
		user          *user.User
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name: "actualiza usuario exitosamente",
			user: &user.User{
				ID:        1,
				Name:      &name,
				Email:     "updated@example.com",
				Password:  "hash",
				RoleID:    1,
				CreatedAt: time.Now(),
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "user" SET`)).
					WithArgs(name, "updated@example.com", "hash", 1, sqlmock.AnyArg(), 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name: "retorna error cuando update falla",
			user: &user.User{
				ID:       1,
				Name:     &name,
				Email:    "updated@example.com",
				Password: "hash",
				RoleID:   1,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "user" SET`)).
					WithArgs(name, "updated@example.com", "hash", 1, sqlmock.AnyArg(), 1).
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

			err := repo.Update(ctx, tt.user)

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
		userID        int
		mockSetup     func(sqlmock.Sqlmock)
		expectedError bool
	}{
		{
			name:   "elimina usuario exitosamente",
			userID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "user" WHERE "user"."id" = $1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			expectedError: false,
		},
		{
			name:   "retorna error cuando delete falla",
			userID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "user" WHERE "user"."id" = $1`)).
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

			err := repo.Delete(ctx, tt.userID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

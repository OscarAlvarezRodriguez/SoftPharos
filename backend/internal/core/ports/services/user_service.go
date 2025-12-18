package services

import (
	"context"
	"softpharos/internal/core/domain/user"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]user.User, error)
	GetUserByID(ctx context.Context, id int) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	CreateUser(ctx context.Context, user *user.User) error
	UpdateUser(ctx context.Context, user *user.User) error
	DeleteUser(ctx context.Context, id int) error
}

package repository

import (
	"context"
	"softpharos/internal/core/domain/user"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]user.User, error)
	GetByID(ctx context.Context, id int) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, id int) error
}

package services

import (
	"context"
	"softpharos/internal/core/domain/user"
)

type AuthService interface {
	AuthenticateWithGoogle(ctx context.Context, idToken string) (*user.User, string, error)
}

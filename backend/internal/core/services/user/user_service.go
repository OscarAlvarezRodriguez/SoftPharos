package user

import (
	"context"
	"softpharos/internal/core/domain/user"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	userRepo repository.UserRepository
}

func New(userRepo repository.UserRepository) services.UserService {
	return &Service{
		userRepo: userRepo,
	}
}

func (s *Service) GetAllUsers(ctx context.Context) ([]user.User, error) {
	return s.userRepo.GetAll(ctx)
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*user.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

func (s *Service) CreateUser(ctx context.Context, usr *user.User) error {
	return s.userRepo.Create(ctx, usr)
}

func (s *Service) UpdateUser(ctx context.Context, usr *user.User) error {
	return s.userRepo.Update(ctx, usr)
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	return s.userRepo.Delete(ctx, id)
}

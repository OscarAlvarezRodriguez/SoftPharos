package auth

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"softpharos/internal/auth"
	"softpharos/internal/core/domain/user"
	"softpharos/internal/core/ports/repository"
	"softpharos/internal/core/ports/services"
)

type Service struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func New(userRepo repository.UserRepository, roleRepo repository.RoleRepository) services.AuthService {
	return &Service{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *Service) AuthenticateWithGoogle(ctx context.Context, idToken string) (*user.User, string, error) {
	tokenInfo, err := auth.VerifyGoogleToken(ctx, idToken)
	if err != nil {
		return nil, "", err
	}

	existingUser, err := s.userRepo.GetByProviderID(ctx, tokenInfo.Sub)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", err
	}

	var domainUser *user.User

	if errors.Is(err, gorm.ErrRecordNotFound) {
		studentRole, err := s.roleRepo.GetByName(ctx, "student")
		if err != nil {
			return nil, "", err
		}

		domainUser = &user.User{
			Name:       &tokenInfo.Name,
			Email:      tokenInfo.Email,
			ProviderID: tokenInfo.Sub,
			RoleID:     studentRole.ID,
			PictureURL: &tokenInfo.Picture,
		}

		if err := s.userRepo.Create(ctx, domainUser); err != nil {
			return nil, "", err
		}

		domainUser, err = s.userRepo.GetByID(ctx, domainUser.ID)
		if err != nil {
			return nil, "", err
		}
	} else {
		domainUser = existingUser

		updated := false
		if tokenInfo.Name != "" && (domainUser.Name == nil || *domainUser.Name != tokenInfo.Name) {
			domainUser.Name = &tokenInfo.Name
			updated = true
		}
		if tokenInfo.Picture != "" && (domainUser.PictureURL == nil || *domainUser.PictureURL != tokenInfo.Picture) {
			domainUser.PictureURL = &tokenInfo.Picture
			updated = true
		}

		if updated {
			_ = s.userRepo.Update(ctx, domainUser)
		}
	}

	accessToken, err := auth.GenerateJWT(domainUser.ID, domainUser.Email, domainUser.RoleID)
	if err != nil {
		return nil, "", err
	}

	return domainUser, accessToken, nil
}

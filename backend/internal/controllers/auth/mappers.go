package auth

import (
	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/domain/user"
)

func ToUserInfo(u *user.User, accessToken string) *AuthResponse {
	if u == nil {
		return nil
	}

	roleName := "student"
	if u.Role != nil {
		roleName = u.Role.Name
	}

	return &AuthResponse{
		AccessToken: accessToken,
		User: UserInfo{
			ID:         u.ID,
			Name:       u.Name,
			Email:      u.Email,
			RoleID:     u.RoleID,
			RoleName:   roleName,
			PictureURL: u.PictureURL,
			CreatedAt:  u.CreatedAt,
		},
	}
}

func ToRoleResponse(r *role.Role) *RoleInfo {
	if r == nil {
		return nil
	}

	return &RoleInfo{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

type RoleInfo struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

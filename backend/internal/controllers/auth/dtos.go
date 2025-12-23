package auth

import "time"

type GoogleLoginRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}

type AuthResponse struct {
	AccessToken string   `json:"accessToken"`
	User        UserInfo `json:"user"`
}

type UserInfo struct {
	ID         int       `json:"id"`
	Name       *string   `json:"name"`
	Email      string    `json:"email"`
	RoleID     int       `json:"roleId"`
	RoleName   string    `json:"roleName"`
	PictureURL *string   `json:"pictureUrl,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

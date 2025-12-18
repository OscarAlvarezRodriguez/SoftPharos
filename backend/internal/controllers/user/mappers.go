package user

import (
	"softpharos/internal/core/domain/role"
	"softpharos/internal/core/domain/user"
)

func ToUserDomain(req *CreateUserRequest) *user.User {
	return &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		RoleID:   req.RoleID,
	}
}

func ToUserResponse(usr *user.User) *UserResponse {
	if usr == nil {
		return nil
	}

	response := &UserResponse{
		ID:        usr.ID,
		Name:      usr.Name,
		Email:     usr.Email,
		RoleID:    usr.RoleID,
		CreatedAt: usr.CreatedAt,
	}

	if usr.Role != nil {
		response.Role = ToRoleResponse(usr.Role)
	}

	return response
}

func ToRoleResponse(r *role.Role) *RoleResponse {
	if r == nil {
		return nil
	}

	return &RoleResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

func ToUserListResponse(users []user.User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, usr := range users {
		responses[i] = *ToUserResponse(&usr)
	}
	return responses
}

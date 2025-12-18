package role

import (
	"softpharos/internal/core/domain/role"
)

func ToRoleResponse(r *role.Role) *RoleResponse {
	if r == nil {
		return nil
	}

	return &RoleResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
	}
}

func ToRoleListResponse(roles []role.Role) []RoleResponse {
	responses := make([]RoleResponse, len(roles))
	for i, r := range roles {
		responses[i] = *ToRoleResponse(&r)
	}
	return responses
}

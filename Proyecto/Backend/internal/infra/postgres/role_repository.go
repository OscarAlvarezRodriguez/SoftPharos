package postgres

import (
	"context"
	"softpharos/internal/core/domain"
)

// RoleRepository maneja las operaciones de la entidad Role
type RoleRepository struct {
	client *Client
}

// NewRoleRepository crea una nueva instancia del repositorio de roles
func NewRoleRepository(client *Client) *RoleRepository {
	return &RoleRepository{client: client}
}

// GetAll obtiene todos los roles
func (r *RoleRepository) GetAll(ctx context.Context) ([]domain.Role, error) {
	var roles []domain.Role
	result := r.client.DB.WithContext(ctx).Find(&roles)
	return roles, result.Error
}

// GetByID obtiene un rol por su ID
func (r *RoleRepository) GetByID(ctx context.Context, id int) (*domain.Role, error) {
	var role domain.Role
	result := r.client.DB.WithContext(ctx).First(&role, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

// GetByName obtiene un rol por su nombre
func (r *RoleRepository) GetByName(ctx context.Context, name string) (*domain.Role, error) {
	var role domain.Role
	result := r.client.DB.WithContext(ctx).Where("name = ?", name).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

// Create crea un nuevo rol
func (r *RoleRepository) Create(ctx context.Context, role *domain.Role) error {
	return r.client.DB.WithContext(ctx).Create(role).Error
}

// Update actualiza un rol existente
func (r *RoleRepository) Update(ctx context.Context, role *domain.Role) error {
	return r.client.DB.WithContext(ctx).Save(role).Error
}

// Delete elimina un rol
func (r *RoleRepository) Delete(ctx context.Context, id int) error {
	return r.client.DB.WithContext(ctx).Delete(&domain.Role{}, id).Error
}

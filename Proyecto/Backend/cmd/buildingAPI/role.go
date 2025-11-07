package buildingAPI

import (
	roleController "softpharos/internal/controllers/role"
	roleRepo "softpharos/internal/core/repository/role"
	"softpharos/internal/core/services/role"
	"softpharos/internal/infra/databases"
)

// BuildRoleController construye y retorna el contenedor con todas las dependencias
func BuildRoleController(dbClient *databases.Client) *roleController.Controller {
	roleRepository := roleRepo.New(dbClient)
	roleService := role.New(roleRepository)
	roleCtrl := roleController.New(roleService)

	return roleCtrl
}

package buildingAPI

import (
	"github.com/gin-gonic/gin"
	roleController "softpharos/internal/controllers/role"
	roleRepo "softpharos/internal/core/repository/role"
	"softpharos/internal/core/services/role"
	"softpharos/internal/infra/databases"
)

func BuildRoleController() *roleController.Controller {
	dbClient := databases.GetInstance()
	roleRepository := roleRepo.New(dbClient)
	roleService := role.New(roleRepository)
	roleCtrl := roleController.New(roleService)

	return roleCtrl
}

func RegisterRoleRoutes(router *gin.RouterGroup) {
	roleCtrl := BuildRoleController()

	roles := router.Group("/roles")
	{
		roles.GET("", roleCtrl.GetAllRoles)
		roles.GET("/:id", roleCtrl.GetRoleByID)
		roles.GET("/name/:name", roleCtrl.GetRoleByName)
	}
}

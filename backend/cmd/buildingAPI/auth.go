package buildingAPI

import (
	"github.com/gin-gonic/gin"

	authController "softpharos/internal/controllers/auth"
	roleRepo "softpharos/internal/core/repository/role"
	userRepo "softpharos/internal/core/repository/user"
	authService "softpharos/internal/core/services/auth"
	"softpharos/internal/infra/databases"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	client := databases.GetInstance()

	userRepository := userRepo.New(client)
	roleRepository := roleRepo.New(client)
	service := authService.New(userRepository, roleRepository)
	controller := authController.New(service)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/google", controller.GoogleLogin)
	}
}

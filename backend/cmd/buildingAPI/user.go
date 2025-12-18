package buildingAPI

import (
	"github.com/gin-gonic/gin"
	userController "softpharos/internal/controllers/user"
	userRepo "softpharos/internal/core/repository/user"
	"softpharos/internal/core/services/user"
	"softpharos/internal/infra/databases"
)

func BuildUserController() *userController.Controller {
	dbClient := databases.GetInstance()
	repo := userRepo.New(dbClient)
	service := user.New(repo)
	ctrl := userController.New(service)

	return ctrl
}

func RegisterUserRoutes(router *gin.RouterGroup) {
	userCtrl := BuildUserController()

	users := router.Group("/users")
	{
		users.GET("", userCtrl.GetAllUsers)
		users.GET("/:id", userCtrl.GetUserByID)
		users.GET("/email/:email", userCtrl.GetUserByEmail)
		users.POST("", userCtrl.CreateUser)
		users.PUT("/:id", userCtrl.UpdateUser)
		users.DELETE("/:id", userCtrl.DeleteUser)
	}
}

package app

import (
	"github.com/gin-gonic/gin"

	"softpharos/cmd/buildingAPI"
)

func MapUrls(router *gin.Engine) {
	v1 := router.Group("")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "SoftPharos API is running",
			})
		})

		// Registrar rutas de cada dominio
		buildingAPI.RegisterProjectRoutes(v1)
		buildingAPI.RegisterRoleRoutes(v1)
		buildingAPI.RegisterUserRoutes(v1)
		buildingAPI.RegisterMilestoneRoutes(v1)
		buildingAPI.RegisterCommentRoutes(v1)
		buildingAPI.RegisterDeliverableRoutes(v1)
		buildingAPI.RegisterFeedbackRoutes(v1)
		buildingAPI.RegisterProjectMemberRoutes(v1)
		buildingAPI.RegisterReactionRoutes(v1)
	}
}

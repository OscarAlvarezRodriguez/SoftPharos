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
	}
}

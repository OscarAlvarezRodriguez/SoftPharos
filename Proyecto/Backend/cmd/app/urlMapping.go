package app

import (
	projectController "softpharos/internal/controllers/project"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, projectCtrl *projectController.Controller) {
	v1 := router.Group("")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "SoftPharos API is running",
			})
		})

		// Rutas de proyectos
		projects := v1.Group("/projects")
		{
			projects.GET("", projectCtrl.GetAllProjects)
			projects.GET("/:id", projectCtrl.GetProjectByID)
			projects.GET("/owner/:owner", projectCtrl.GetProjectsByOwner)
			projects.POST("", projectCtrl.CreateProject)
			projects.PUT("/:id", projectCtrl.UpdateProject)
			projects.DELETE("/:id", projectCtrl.DeleteProject)
		}
	}
}

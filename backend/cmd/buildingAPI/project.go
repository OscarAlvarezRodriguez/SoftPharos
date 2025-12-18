package buildingAPI

import (
	"github.com/gin-gonic/gin"
	projectController "softpharos/internal/controllers/project"
	project2 "softpharos/internal/core/repository/project"
	"softpharos/internal/core/services/project"
	"softpharos/internal/infra/databases"
)

func BuildProjectController() *projectController.Controller {
	dbClient := databases.GetInstance()
	projectRepo := project2.New(dbClient)
	projectService := project.New(projectRepo)
	projectCtrl := projectController.New(projectService)

	return projectCtrl
}

func RegisterProjectRoutes(router *gin.RouterGroup) {
	projectCtrl := BuildProjectController()

	projects := router.Group("/projects")
	{
		projects.GET("", projectCtrl.GetAllProjects)
		projects.GET("/:id", projectCtrl.GetProjectByID)
		projects.GET("/owner/:owner", projectCtrl.GetProjectsByOwner)
		projects.POST("", projectCtrl.CreateProject)
		projects.PUT("/:id", projectCtrl.UpdateProject)
		projects.DELETE("/:id", projectCtrl.DeleteProject)
	}
}

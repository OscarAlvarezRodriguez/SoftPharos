package buildingAPI

import (
	projectController "softpharos/internal/controllers/project"
	project2 "softpharos/internal/core/repository/project"
	"softpharos/internal/core/services/project"
	"softpharos/internal/infra/databases"
)

// BuildProjectController construye y retorna el contenedor con todas las dependencias
func BuildProjectController(dbClient *databases.Client) *projectController.Controller {
	projectRepo := project2.New(dbClient)
	projectService := project.New(projectRepo)
	projectCtrl := projectController.New(projectService)

	return projectCtrl
}

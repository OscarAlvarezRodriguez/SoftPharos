package buildingAPI

import (
	projectController "softpharos/internal/controllers/project"
	"softpharos/internal/core/repository"
	"softpharos/internal/core/services"
	"softpharos/internal/infra/databases"
)

// BuildProjectController construye y retorna el contenedor con todas las dependencias
func BuildProjectController(dbClient *databases.Client) *projectController.Controller {
	projectRepo := repository.New(dbClient)
	projectService := services.New(projectRepo)
	projectCtrl := projectController.New(projectService)

	return projectCtrl
}

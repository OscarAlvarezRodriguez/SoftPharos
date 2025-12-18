package project

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	projectService services.ProjectService
}

func New(projectService services.ProjectService) *Controller {
	return &Controller{
		projectService: projectService,
	}
}

func (c *Controller) GetAllProjects(ctx *gin.Context) {
	projects, err := c.projectService.GetAllProjects(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectListResponse(projects))
}

func (c *Controller) GetProjectByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	project, err := c.projectService.GetProjectByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Proyecto no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectResponse(project))
}

func (c *Controller) GetProjectsByOwner(ctx *gin.Context) {
	OwnerIDParam := ctx.Param("ownerId")
	OwnerId, err := strconv.Atoi(OwnerIDParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID del creador debe ser un número válido")
		return
	}

	projects, err := c.projectService.GetProjectsByOwner(ctx.Request.Context(), OwnerId)
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectListResponse(projects))
}

func (c *Controller) CreateProject(ctx *gin.Context) {
	var req CreateProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	project := ToProjectDomain(&req)
	if err := c.projectService.CreateProject(ctx.Request.Context(), project); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToProjectResponse(project))
}

func (c *Controller) UpdateProject(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	// Primero obtenemos el proyecto existente
	existingProject, err := c.projectService.GetProjectByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Proyecto no encontrado")
		return
	}

	// Actualizamos solo los campos proporcionados
	if req.Name != nil {
		existingProject.Name = req.Name
	}
	if req.Objective != nil {
		existingProject.Objective = req.Objective
	}

	if err := c.projectService.UpdateProject(ctx.Request.Context(), existingProject); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectResponse(existingProject))
}

func (c *Controller) DeleteProject(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.projectService.DeleteProject(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Proyecto eliminado exitosamente",
	})
}

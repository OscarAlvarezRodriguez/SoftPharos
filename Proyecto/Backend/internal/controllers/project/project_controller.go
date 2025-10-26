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
		ctx.JSON(http.StatusInternalServerError, controllers.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ToProjectListResponse(projects))
}

func (c *Controller) GetProjectByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "El ID debe ser un número válido",
		})
		return
	}

	project, err := c.projectService.GetProjectByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Proyecto no encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, ToProjectResponse(project))
}

func (c *Controller) GetProjectsByOwner(ctx *gin.Context) {
	creatorIDParam := ctx.Param("creatorId")
	creatorID, err := strconv.Atoi(creatorIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "El ID del creador debe ser un número válido",
		})
		return
	}

	projects, err := c.projectService.GetProjectsByCreator(ctx.Request.Context(), creatorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, controllers.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ToProjectListResponse(projects))
}

func (c *Controller) CreateProject(ctx *gin.Context) {
	var req CreateProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: err.Error(),
		})
		return
	}

	project := ToProjectDomain(&req)
	if err := c.projectService.CreateProject(ctx.Request.Context(), project); err != nil {
		ctx.JSON(http.StatusInternalServerError, controllers.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, ToProjectResponse(project))
}

func (c *Controller) UpdateProject(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "El ID debe ser un número válido",
		})
		return
	}

	var req UpdateProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_REQUEST",
			Message: err.Error(),
		})
		return
	}

	// Primero obtenemos el proyecto existente
	existingProject, err := c.projectService.GetProjectByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Proyecto no encontrado",
		})
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
		ctx.JSON(http.StatusInternalServerError, controllers.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ToProjectResponse(existingProject))
}

func (c *Controller) DeleteProject(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "El ID debe ser un número válido",
		})
		return
	}

	if err := c.projectService.DeleteProject(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, controllers.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.SuccessResponse{
		Message: "Proyecto eliminado exitosamente",
	})
}

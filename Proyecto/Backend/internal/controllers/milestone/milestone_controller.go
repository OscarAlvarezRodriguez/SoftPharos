package milestone

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	milestoneService services.MilestoneService
}

func New(milestoneService services.MilestoneService) *Controller {
	return &Controller{
		milestoneService: milestoneService,
	}
}

func (c *Controller) GetAllMilestones(ctx *gin.Context) {
	milestones, err := c.milestoneService.GetAllMilestones(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToMilestoneListResponse(milestones))
}

func (c *Controller) GetMilestoneByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	milestone, err := c.milestoneService.GetMilestoneByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Milestone no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToMilestoneResponse(milestone))
}

func (c *Controller) GetMilestonesByProjectID(ctx *gin.Context) {
	projectIDParam := ctx.Param("projectId")
	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID del proyecto debe ser un número válido")
		return
	}

	milestones, err := c.milestoneService.GetMilestonesByProjectID(ctx.Request.Context(), projectID)
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToMilestoneListResponse(milestones))
}

func (c *Controller) CreateMilestone(ctx *gin.Context) {
	var req CreateMilestoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	milestone := ToMilestoneDomain(&req)
	if err := c.milestoneService.CreateMilestone(ctx.Request.Context(), milestone); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToMilestoneResponse(milestone))
}

func (c *Controller) UpdateMilestone(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateMilestoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	existingMilestone, err := c.milestoneService.GetMilestoneByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Milestone no encontrado")
		return
	}

	if req.Title != nil {
		existingMilestone.Title = req.Title
	}
	if req.Description != nil {
		existingMilestone.Description = req.Description
	}
	if req.ClassWeek != nil {
		existingMilestone.ClassWeek = req.ClassWeek
	}

	if err := c.milestoneService.UpdateMilestone(ctx.Request.Context(), existingMilestone); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToMilestoneResponse(existingMilestone))
}

func (c *Controller) DeleteMilestone(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.milestoneService.DeleteMilestone(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Milestone eliminado exitosamente",
	})
}

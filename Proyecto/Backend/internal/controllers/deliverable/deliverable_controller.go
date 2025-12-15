package deliverable

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	deliverableService services.DeliverableService
}

func New(deliverableService services.DeliverableService) *Controller {
	return &Controller{
		deliverableService: deliverableService,
	}
}

func (c *Controller) GetAllDeliverables(ctx *gin.Context) {
	deliverables, err := c.deliverableService.GetAllDeliverables(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToDeliverableListResponse(deliverables))
}

func (c *Controller) GetDeliverableByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	deliverable, err := c.deliverableService.GetDeliverableByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Entregable no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToDeliverableResponse(deliverable))
}

func (c *Controller) GetDeliverablesByMilestoneID(ctx *gin.Context) {
	milestoneIDParam := ctx.Param("milestoneId")
	milestoneID, err := strconv.Atoi(milestoneIDParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID del milestone debe ser un número válido")
		return
	}

	deliverables, err := c.deliverableService.GetDeliverablesByMilestoneID(ctx.Request.Context(), milestoneID)
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToDeliverableListResponse(deliverables))
}

func (c *Controller) CreateDeliverable(ctx *gin.Context) {
	var req CreateDeliverableRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	deliverable := ToDeliverableDomain(&req)
	if err := c.deliverableService.CreateDeliverable(ctx.Request.Context(), deliverable); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToDeliverableResponse(deliverable))
}

func (c *Controller) UpdateDeliverable(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateDeliverableRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	existingDeliverable, err := c.deliverableService.GetDeliverableByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Entregable no encontrado")
		return
	}

	if req.URL != nil {
		existingDeliverable.URL = *req.URL
	}
	if req.Type != nil {
		existingDeliverable.Type = req.Type
	}

	if err := c.deliverableService.UpdateDeliverable(ctx.Request.Context(), existingDeliverable); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToDeliverableResponse(existingDeliverable))
}

func (c *Controller) DeleteDeliverable(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.deliverableService.DeleteDeliverable(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Entregable eliminado exitosamente",
	})
}

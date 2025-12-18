package reaction

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	reactionService services.ReactionService
}

func New(reactionService services.ReactionService) *Controller {
	return &Controller{
		reactionService: reactionService,
	}
}

func (c *Controller) GetAllReactions(ctx *gin.Context) {
	reactions, err := c.reactionService.GetAllReactions(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToReactionListResponse(reactions))
}

func (c *Controller) GetReactionByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	reaction, err := c.reactionService.GetReactionByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Reacción no encontrada")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToReactionResponse(reaction))
}

func (c *Controller) GetReactionsByMilestoneID(ctx *gin.Context) {
	milestoneIDParam := ctx.Param("milestoneId")
	milestoneID, err := strconv.Atoi(milestoneIDParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID del milestone debe ser un número válido")
		return
	}

	reactions, err := c.reactionService.GetReactionsByMilestoneID(ctx.Request.Context(), milestoneID)
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToReactionListResponse(reactions))
}

func (c *Controller) CreateReaction(ctx *gin.Context) {
	var req CreateReactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	reaction := ToReactionDomain(&req)
	if err := c.reactionService.CreateReaction(ctx.Request.Context(), reaction); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToReactionResponse(reaction))
}

func (c *Controller) UpdateReaction(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateReactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	existingReaction, err := c.reactionService.GetReactionByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Reacción no encontrada")
		return
	}

	if req.Type != nil {
		existingReaction.Type = req.Type
	}

	if err := c.reactionService.UpdateReaction(ctx.Request.Context(), existingReaction); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToReactionResponse(existingReaction))
}

func (c *Controller) DeleteReaction(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.reactionService.DeleteReaction(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Reacción eliminada exitosamente",
	})
}

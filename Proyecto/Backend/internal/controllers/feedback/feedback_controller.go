package feedback

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	feedbackService services.FeedbackService
}

func New(feedbackService services.FeedbackService) *Controller {
	return &Controller{
		feedbackService: feedbackService,
	}
}

func (c *Controller) GetAllFeedbacks(ctx *gin.Context) {
	feedbacks, err := c.feedbackService.GetAllFeedbacks(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToFeedbackListResponse(feedbacks))
}

func (c *Controller) GetFeedbackByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	feedback, err := c.feedbackService.GetFeedbackByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Feedback no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToFeedbackResponse(feedback))
}

func (c *Controller) GetFeedbacksByMilestoneID(ctx *gin.Context) {
	milestoneIDParam := ctx.Param("milestoneId")
	milestoneID, err := strconv.Atoi(milestoneIDParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID del milestone debe ser un número válido")
		return
	}

	feedbacks, err := c.feedbackService.GetFeedbacksByMilestoneID(ctx.Request.Context(), milestoneID)
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToFeedbackListResponse(feedbacks))
}

func (c *Controller) CreateFeedback(ctx *gin.Context) {
	var req CreateFeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	feedback := ToFeedbackDomain(&req)
	if err := c.feedbackService.CreateFeedback(ctx.Request.Context(), feedback); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToFeedbackResponse(feedback))
}

func (c *Controller) UpdateFeedback(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateFeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	existingFeedback, err := c.feedbackService.GetFeedbackByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Feedback no encontrado")
		return
	}

	existingFeedback.Content = req.Content

	if err := c.feedbackService.UpdateFeedback(ctx.Request.Context(), existingFeedback); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToFeedbackResponse(existingFeedback))
}

func (c *Controller) DeleteFeedback(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.feedbackService.DeleteFeedback(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Feedback eliminado exitosamente",
	})
}

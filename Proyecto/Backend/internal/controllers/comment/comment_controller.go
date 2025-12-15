package comment

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	commentService services.CommentService
}

func New(commentService services.CommentService) *Controller {
	return &Controller{
		commentService: commentService,
	}
}

func (c *Controller) GetAllComments(ctx *gin.Context) {
	comments, err := c.commentService.GetAllComments(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToCommentListResponse(comments))
}

func (c *Controller) GetCommentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	comment, err := c.commentService.GetCommentByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Comentario no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToCommentResponse(comment))
}

func (c *Controller) GetCommentsByMilestoneID(ctx *gin.Context) {
	milestoneIDParam := ctx.Param("milestoneId")
	milestoneID, err := strconv.Atoi(milestoneIDParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID del milestone debe ser un número válido")
		return
	}

	comments, err := c.commentService.GetCommentsByMilestoneID(ctx.Request.Context(), milestoneID)
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToCommentListResponse(comments))
}

func (c *Controller) CreateComment(ctx *gin.Context) {
	var req CreateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	comment := ToCommentDomain(&req)
	if err := c.commentService.CreateComment(ctx.Request.Context(), comment); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToCommentResponse(comment))
}

func (c *Controller) UpdateComment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	existingComment, err := c.commentService.GetCommentByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Comentario no encontrado")
		return
	}

	if req.Content != nil {
		existingComment.Content = req.Content
	}

	if err := c.commentService.UpdateComment(ctx.Request.Context(), existingComment); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToCommentResponse(existingComment))
}

func (c *Controller) DeleteComment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.commentService.DeleteComment(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Comentario eliminado exitosamente",
	})
}

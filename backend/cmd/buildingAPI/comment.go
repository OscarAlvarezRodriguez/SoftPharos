package buildingAPI

import (
	"github.com/gin-gonic/gin"
	commentController "softpharos/internal/controllers/comment"
	commentRepo "softpharos/internal/core/repository/comment"
	"softpharos/internal/core/services/comment"
	"softpharos/internal/infra/databases"
)

func BuildCommentController() *commentController.Controller {
	dbClient := databases.GetInstance()
	repo := commentRepo.New(dbClient)
	service := comment.New(repo)
	ctrl := commentController.New(service)

	return ctrl
}

func RegisterCommentRoutes(router *gin.RouterGroup) {
	commentCtrl := BuildCommentController()

	comments := router.Group("/comments")
	{
		comments.GET("", commentCtrl.GetAllComments)
		comments.GET("/:id", commentCtrl.GetCommentByID)
		comments.GET("/milestone/:milestoneId", commentCtrl.GetCommentsByMilestoneID)
		comments.POST("", commentCtrl.CreateComment)
		comments.PUT("/:id", commentCtrl.UpdateComment)
		comments.DELETE("/:id", commentCtrl.DeleteComment)
	}
}

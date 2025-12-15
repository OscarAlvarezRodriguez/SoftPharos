package buildingAPI

import (
	"github.com/gin-gonic/gin"
	reactionController "softpharos/internal/controllers/reaction"
	reactionRepo "softpharos/internal/core/repository/reaction"
	"softpharos/internal/core/services/reaction"
	"softpharos/internal/infra/databases"
)

func BuildReactionController() *reactionController.Controller {
	dbClient := databases.GetInstance()
	repo := reactionRepo.New(dbClient)
	service := reaction.New(repo)
	ctrl := reactionController.New(service)

	return ctrl
}

func RegisterReactionRoutes(router *gin.RouterGroup) {
	reactionCtrl := BuildReactionController()

	reactions := router.Group("/reactions")
	{
		reactions.GET("", reactionCtrl.GetAllReactions)
		reactions.GET("/:id", reactionCtrl.GetReactionByID)
		reactions.GET("/milestone/:milestoneId", reactionCtrl.GetReactionsByMilestoneID)
		reactions.POST("", reactionCtrl.CreateReaction)
		reactions.PUT("/:id", reactionCtrl.UpdateReaction)
		reactions.DELETE("/:id", reactionCtrl.DeleteReaction)
	}
}

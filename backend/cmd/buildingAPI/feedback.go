package buildingAPI

import (
	"github.com/gin-gonic/gin"
	feedbackController "softpharos/internal/controllers/feedback"
	feedbackRepo "softpharos/internal/core/repository/feedback"
	"softpharos/internal/core/services/feedback"
	"softpharos/internal/infra/databases"
)

func BuildFeedbackController() *feedbackController.Controller {
	dbClient := databases.GetInstance()
	repo := feedbackRepo.New(dbClient)
	service := feedback.New(repo)
	ctrl := feedbackController.New(service)

	return ctrl
}

func RegisterFeedbackRoutes(router *gin.RouterGroup) {
	feedbackCtrl := BuildFeedbackController()

	feedbacks := router.Group("/feedbacks")
	{
		feedbacks.GET("", feedbackCtrl.GetAllFeedbacks)
		feedbacks.GET("/:id", feedbackCtrl.GetFeedbackByID)
		feedbacks.GET("/milestone/:milestoneId", feedbackCtrl.GetFeedbacksByMilestoneID)
		feedbacks.POST("", feedbackCtrl.CreateFeedback)
		feedbacks.PUT("/:id", feedbackCtrl.UpdateFeedback)
		feedbacks.DELETE("/:id", feedbackCtrl.DeleteFeedback)
	}
}

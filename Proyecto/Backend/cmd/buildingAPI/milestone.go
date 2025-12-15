package buildingAPI

import (
	"github.com/gin-gonic/gin"
	milestoneController "softpharos/internal/controllers/milestone"
	milestoneRepo "softpharos/internal/core/repository/milestone"
	"softpharos/internal/core/services/milestone"
	"softpharos/internal/infra/databases"
)

func BuildMilestoneController() *milestoneController.Controller {
	dbClient := databases.GetInstance()
	repo := milestoneRepo.New(dbClient)
	service := milestone.New(repo)
	ctrl := milestoneController.New(service)

	return ctrl
}

func RegisterMilestoneRoutes(router *gin.RouterGroup) {
	milestoneCtrl := BuildMilestoneController()

	milestones := router.Group("/milestones")
	{
		milestones.GET("", milestoneCtrl.GetAllMilestones)
		milestones.GET("/:id", milestoneCtrl.GetMilestoneByID)
		milestones.GET("/project/:projectId", milestoneCtrl.GetMilestonesByProjectID)
		milestones.POST("", milestoneCtrl.CreateMilestone)
		milestones.PUT("/:id", milestoneCtrl.UpdateMilestone)
		milestones.DELETE("/:id", milestoneCtrl.DeleteMilestone)
	}
}

package buildingAPI

import (
	"github.com/gin-gonic/gin"
	deliverableController "softpharos/internal/controllers/deliverable"
	deliverableRepo "softpharos/internal/core/repository/deliverable"
	"softpharos/internal/core/services/deliverable"
	"softpharos/internal/infra/databases"
)

func BuildDeliverableController() *deliverableController.Controller {
	dbClient := databases.GetInstance()
	repo := deliverableRepo.New(dbClient)
	service := deliverable.New(repo)
	ctrl := deliverableController.New(service)

	return ctrl
}

func RegisterDeliverableRoutes(router *gin.RouterGroup) {
	deliverableCtrl := BuildDeliverableController()

	deliverables := router.Group("/deliverables")
	{
		deliverables.GET("", deliverableCtrl.GetAllDeliverables)
		deliverables.GET("/:id", deliverableCtrl.GetDeliverableByID)
		deliverables.GET("/milestone/:milestoneId", deliverableCtrl.GetDeliverablesByMilestoneID)
		deliverables.POST("", deliverableCtrl.CreateDeliverable)
		deliverables.PUT("/:id", deliverableCtrl.UpdateDeliverable)
		deliverables.DELETE("/:id", deliverableCtrl.DeleteDeliverable)
	}
}

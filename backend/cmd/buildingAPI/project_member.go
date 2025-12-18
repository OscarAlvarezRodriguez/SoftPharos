package buildingAPI

import (
	"github.com/gin-gonic/gin"
	projectMemberController "softpharos/internal/controllers/project_member"
	projectMemberRepo "softpharos/internal/core/repository/project_member"
	"softpharos/internal/core/services/project_member"
	"softpharos/internal/infra/databases"
)

func BuildProjectMemberController() *projectMemberController.Controller {
	dbClient := databases.GetInstance()
	repo := projectMemberRepo.New(dbClient)
	service := project_member.New(repo)
	ctrl := projectMemberController.New(service)

	return ctrl
}

func RegisterProjectMemberRoutes(router *gin.RouterGroup) {
	projectMemberCtrl := BuildProjectMemberController()

	projectMembers := router.Group("/project-members")
	{
		projectMembers.GET("", projectMemberCtrl.GetAllProjectMembers)
		projectMembers.GET("/:id", projectMemberCtrl.GetProjectMemberByID)
		projectMembers.GET("/project/:projectId", projectMemberCtrl.GetProjectMembersByProjectID)
		projectMembers.POST("", projectMemberCtrl.CreateProjectMember)
		projectMembers.PUT("/:id", projectMemberCtrl.UpdateProjectMember)
		projectMembers.DELETE("/:id", projectMemberCtrl.DeleteProjectMember)
	}
}

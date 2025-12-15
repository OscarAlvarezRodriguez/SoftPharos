package project_member

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	projectMemberService services.ProjectMemberService
}

func New(projectMemberService services.ProjectMemberService) *Controller {
	return &Controller{
		projectMemberService: projectMemberService,
	}
}

func (c *Controller) GetAllProjectMembers(ctx *gin.Context) {
	projectMembers, err := c.projectMemberService.GetAllProjectMembers(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectMemberListResponse(projectMembers))
}

func (c *Controller) GetProjectMemberByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	projectMember, err := c.projectMemberService.GetProjectMemberByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Miembro del proyecto no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectMemberResponse(projectMember))
}

func (c *Controller) GetProjectMembersByProjectID(ctx *gin.Context) {
	projectIDParam := ctx.Param("projectId")
	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID del proyecto debe ser un número válido")
		return
	}

	projectMembers, err := c.projectMemberService.GetProjectMembersByProjectID(ctx.Request.Context(), projectID)
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectMemberListResponse(projectMembers))
}

func (c *Controller) CreateProjectMember(ctx *gin.Context) {
	var req CreateProjectMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	projectMember := ToProjectMemberDomain(&req)
	if err := c.projectMemberService.CreateProjectMember(ctx.Request.Context(), projectMember); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToProjectMemberResponse(projectMember))
}

func (c *Controller) UpdateProjectMember(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateProjectMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	existingProjectMember, err := c.projectMemberService.GetProjectMemberByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Miembro del proyecto no encontrado")
		return
	}

	if req.Role != nil {
		existingProjectMember.Role = req.Role
	}

	if err := c.projectMemberService.UpdateProjectMember(ctx.Request.Context(), existingProjectMember); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToProjectMemberResponse(existingProjectMember))
}

func (c *Controller) DeleteProjectMember(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.projectMemberService.DeleteProjectMember(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Miembro del proyecto eliminado exitosamente",
	})
}

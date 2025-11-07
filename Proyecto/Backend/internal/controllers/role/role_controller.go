package role

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	roleService services.RoleService
}

func New(roleService services.RoleService) *Controller {
	return &Controller{
		roleService: roleService,
	}
}

func (c *Controller) GetAllRoles(ctx *gin.Context) {
	roles, err := c.roleService.GetAllRoles(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, controllers.ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ToRoleListResponse(roles))
}

func (c *Controller) GetRoleByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "El ID debe ser un número válido",
		})
		return
	}

	role, err := c.roleService.GetRoleByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Rol no encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, ToRoleResponse(role))
}

func (c *Controller) GetRoleByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Error:   "INVALID_NAME",
			Message: "El nombre del rol es requerido",
		})
		return
	}

	role, err := c.roleService.GetRoleByName(ctx.Request.Context(), name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Rol no encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, ToRoleResponse(role))
}

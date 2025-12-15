package user

import (
	"net/http"
	"softpharos/internal/controllers"
	"strconv"

	"softpharos/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService services.UserService
}

func New(userService services.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (c *Controller) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers(ctx.Request.Context())
	if err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToUserListResponse(users))
}

func (c *Controller) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	user, err := c.userService.GetUserByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Usuario no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToUserResponse(user))
}

func (c *Controller) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		controllers.Response.BadRequest(ctx, "El email es requerido")
		return
	}

	user, err := c.userService.GetUserByEmail(ctx.Request.Context(), email)
	if err != nil {
		controllers.Response.NotFound(ctx, "Usuario no encontrado")
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToUserResponse(user))
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	user := ToUserDomain(&req)
	if err := c.userService.CreateUser(ctx.Request.Context(), user); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusCreated, ToUserResponse(user))
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, err.Error())
		return
	}

	existingUser, err := c.userService.GetUserByID(ctx.Request.Context(), id)
	if err != nil {
		controllers.Response.NotFound(ctx, "Usuario no encontrado")
		return
	}

	if req.Name != nil {
		existingUser.Name = req.Name
	}
	if req.Password != nil {
		existingUser.Password = *req.Password
	}
	if req.RoleID != nil {
		existingUser.RoleID = *req.RoleID
	}

	if err := c.userService.UpdateUser(ctx.Request.Context(), existingUser); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToUserResponse(existingUser))
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		controllers.Response.InvalidID(ctx, "El ID debe ser un número válido")
		return
	}

	if err := c.userService.DeleteUser(ctx.Request.Context(), id); err != nil {
		controllers.Response.InternalError(ctx, err.Error())
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, gin.H{
		"message": "Usuario eliminado exitosamente",
	})
}

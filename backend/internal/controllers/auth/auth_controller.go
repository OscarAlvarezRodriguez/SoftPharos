package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"softpharos/internal/controllers"
	"softpharos/internal/core/ports/services"
)

type Controller struct {
	authService services.AuthService
}

func New(authService services.AuthService) *Controller {
	return &Controller{
		authService: authService,
	}
}

func (c *Controller) GoogleLogin(ctx *gin.Context) {
	var req GoogleLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controllers.Response.BadRequest(ctx, "idToken es requerido")
		return
	}

	user, accessToken, err := c.authService.AuthenticateWithGoogle(ctx.Request.Context(), req.IDToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token de Google inválido: " + err.Error(),
		})
		return
	}

	controllers.Response.Success(ctx, http.StatusOK, ToUserInfo(user, accessToken))
}

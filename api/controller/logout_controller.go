package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-backend-pos/bootstrap"
	"go-backend-pos/domain"
)

type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
	Env           *bootstrap.Env
}

func (u *LogoutController) Logout(c *gin.Context) {
	accessToken := c.GetString("x-auth-token")
	userId := c.GetString("x-user-id")

	isSuccess, err := u.LogoutUsecase.DeleteAccessToken(c, accessToken, userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if !isSuccess {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "You have been logged out successfully.",
	})
}

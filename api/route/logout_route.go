package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"go-backend-pos/api/controller"
	"go-backend-pos/bootstrap"
	"go-backend-pos/domain"
	"go-backend-pos/mongo"
	"go-backend-pos/repository"
	"go-backend-pos/usecase"
)

func NewLogoutRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	t := repository.NewTokenRepository(db, domain.CollectionToken)
	lc := &controller.LogoutController{
		LogoutUsecase: usecase.NewLogoutUsecase(t, timeout),
		Env:           env,
	}
	group.POST("/logout", lc.Logout)
}

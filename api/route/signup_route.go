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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	t := repository.NewTokenRepository(db, domain.CollectionToken)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, t, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}

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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}

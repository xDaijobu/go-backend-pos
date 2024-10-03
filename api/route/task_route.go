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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}

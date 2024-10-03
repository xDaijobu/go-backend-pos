package route

import (
	"github.com/gin-gonic/gin"
	"go-backend-pos/api/controller"
	"go-backend-pos/bootstrap"
	"go-backend-pos/domain"
	"go-backend-pos/mongo"
	"go-backend-pos/repository"
	"go-backend-pos/usecase"
	"time"
)

func NewCategoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewCategoryRepository(db, domain.CollectionCategory)

	tc := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(cr, timeout),
	}

	group.POST("/category", tc.Create)
	group.GET("/category", tc.Fetch)
}

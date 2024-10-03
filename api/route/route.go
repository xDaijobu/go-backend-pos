package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"go-backend-pos/api/middleware"
	"go-backend-pos/bootstrap"
	"go-backend-pos/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
	NewLogoutRouter(env, timeout, db, protectedRouter)
	NewCategoryRouter(env, timeout, db, protectedRouter)
}

package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewSignupRouter(timeout, db, publicRouter)
	NewLoginRouter(timeout, db, publicRouter)
	NewGetProfileRouter(timeout, publicRouter)
	// NewRefreshTokenRouter(timeout, db, publicRouter)

	// protectedRouter := gin.Group("")
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// NewProfileRouter(env, timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)
}

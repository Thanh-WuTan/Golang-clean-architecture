package route

import (
	"onlyfounds/api/controller"
	"onlyfounds/api/middleware"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func NewGetProfileRouter(timeout time.Duration, group *gin.RouterGroup) {
	cl := controller.GetProfileController{}
	group.GET("/getprofile", middleware.JwtAuthMiddleware(os.Getenv("ACCESS_TOKEN_SECRET")), cl.Profile)
}

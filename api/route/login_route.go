package route

import (
	"onlyfounds/api/controller"
	"onlyfounds/module/user/repository"
	"onlyfounds/module/user/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
	}
	group.POST("/login", lc.Login)
}

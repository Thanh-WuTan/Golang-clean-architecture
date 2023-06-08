package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"onlyfounds/api/controller"
	"onlyfounds/module/user/repository"
	"onlyfounds/module/user/usecase"
)

func NewSignupRouter(timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
	}
	group.POST("/signup", sc.Signup)
}

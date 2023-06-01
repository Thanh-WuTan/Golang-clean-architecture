package controller

import (
	"net/http"
	"onlyfounds/domain"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
}

func (sc *LoginController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}

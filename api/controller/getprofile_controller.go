package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProfileController struct {
}

func (cl *GetProfileController) Profile(c *gin.Context) {
	u := c.MustGet("current_user")
	c.JSON(http.StatusOK, u)
}

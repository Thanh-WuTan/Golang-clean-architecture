package controller

import (
	"fmt"
	"net/http"
	"onlyfounds/common"
	USER_MODEL "onlyfounds/module/user/model"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase USER_MODEL.SignupUsecase
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request USER_MODEL.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	if request.Email[strings.Index(request.Email, "@")+1:] != "vinuni.edu.vn" {
		fmt.Println("!!!!")
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "must use vinuni email"})
		return
	}

	user, err := sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "email is already used"})
		return
	}

	if request.Password != request.ConfirmPassword {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Password and confirm password does not match"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user = USER_MODEL.User{
		UserName: request.Email[:strings.Index(request.Email, "@")],
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully signup!",
	})
}

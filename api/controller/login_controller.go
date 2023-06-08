package controller

import (
	"net/http"
	common "onlyfounds/common"
	USER_MODEL "onlyfounds/module/user/model"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase USER_MODEL.LoginUsecase
}

func (lc *LoginController) Login(c *gin.Context) {
	var request USER_MODEL.LoginRequest
	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("ACCESS_TOKEN_EXPIRY_MINUTE"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, os.Getenv("REFRESH_TOKEN_SECRET"), os.Getenv("REFRESH_TOKEN_EXPIRY_MINUTE"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := USER_MODEL.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.Header("Authorization", "Bearer "+accessToken)
	c.SetCookie("refresh_token", refreshToken, 604800, "/", "", false, true)
	c.JSON(http.StatusOK, loginResponse)
}

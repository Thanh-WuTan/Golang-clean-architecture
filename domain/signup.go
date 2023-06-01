package domain

import (
	"context"
)

type SignupRequest struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirmpassword" binding:"required"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
}

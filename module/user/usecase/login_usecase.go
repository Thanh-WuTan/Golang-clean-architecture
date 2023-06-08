package usecase

import (
	"context"
	"onlyfounds/internal/tokenutil"
	USER_MODEL "onlyfounds/module/user/model"
	"time"
)

type loginUsecase struct {
	userRepository USER_MODEL.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository USER_MODEL.UserRepository, timeout time.Duration) USER_MODEL.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (USER_MODEL.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *USER_MODEL.User, secret string, expiry string) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *USER_MODEL.User, secret string, expiry string) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

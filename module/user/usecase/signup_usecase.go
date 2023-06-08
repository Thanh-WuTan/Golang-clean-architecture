package usecase

import (
	"context"
	"time"

	USER_MODEL "onlyfounds/module/user/model"
)

type signupUsecase struct {
	userRepository USER_MODEL.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository USER_MODEL.UserRepository, timeout time.Duration) USER_MODEL.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *USER_MODEL.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (USER_MODEL.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

package model

import (
	"context"
	common "onlyfounds/common"
)

type User struct {
	common.SQLModel
	UserName string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByEmail(c context.Context, email string) (User, error)
}

package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
}

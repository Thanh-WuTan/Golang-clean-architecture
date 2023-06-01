package domain

import (
	"context"
)

type User struct {
	SQLModel
	UserName string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetByEmail(c context.Context, email string) (User, error)
}

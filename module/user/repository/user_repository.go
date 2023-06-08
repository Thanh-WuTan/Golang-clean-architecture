package repository

import (
	"context"
	USER_MODEL "onlyfounds/module/user/model"

	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) USER_MODEL.UserRepository {
	return &userRepository{database: db}
}

func (ur *userRepository) Create(c context.Context, user *USER_MODEL.User) error {
	result := ur.database.Create(&user)
	return result.Error
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (USER_MODEL.User, error) {
	var user USER_MODEL.User
	result := ur.database.First(&user, "email = ?", email)
	return user, result.Error
}

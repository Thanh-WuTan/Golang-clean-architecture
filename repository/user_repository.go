package repository

import (
	"context"
	"onlyfounds/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{database: db}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	result := ur.database.Create(&user)
	return result.Error
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	result := ur.database.First(&user, "email = ?", email)
	return user, result.Error
}

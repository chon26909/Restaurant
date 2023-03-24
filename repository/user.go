package repository

import (
	"restaurant/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (user *models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (user *models.User, err error) {
	err = r.db.Where("email=?", email).First(&user).Error
	return user, err
}

package repository

import (
	"backend-toko-buku/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) IsEmailExists(email string) (bool, error) {
	var count int64

	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) IsUsernameExists(username string) (bool, error) {
	var count int64

	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

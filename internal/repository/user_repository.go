package repository

import "backend-toko-buku/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	IsEmailExists(email string) (bool, error)
	IsUsernameExists(username string) (bool, error)
}
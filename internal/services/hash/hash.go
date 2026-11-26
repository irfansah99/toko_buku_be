package hash

import (
	"backend-toko-buku/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashed, plain string) error
}

type BcryptHasher struct{}

func (b *BcryptHasher) Hash(password string) (string, error) {
	return utils.HashPassword(password)
}

func (b *BcryptHasher) Compare(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

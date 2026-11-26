package auth

import (
	"backend-toko-buku/internal/dto"
	"backend-toko-buku/internal/models"
	"backend-toko-buku/internal/repository"
	AppErrorpkg "backend-toko-buku/internal/services/errors"
	"backend-toko-buku/internal/services/hash"
	"backend-toko-buku/internal/utils/mappers"
	"fmt"
)

type RegisterService struct {
	userRepo repository.UserRepository
	hasher   hash.PasswordHasher
}

func NewRegisterService(
	userRepo repository.UserRepository,
	hasher hash.PasswordHasher,
) *RegisterService {
	return &RegisterService{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (s *RegisterService) RegisterUser(req dto.UserCreateRequest) (dto.UserResponse, error) {
	errorMap := make(map[string]string)

	EmailExist, err := s.userRepo.IsEmailExists(req.Email)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("find user by email failed: %w", err)
	}
	UsernameExist, err := s.userRepo.IsUsernameExists(req.Username)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("find user by username failed: %w", err)
	}
	if EmailExist {
		errorMap["email"] = "email already exists"
	}
	if UsernameExist {
		errorMap["username"] = "username already exists"
	}
	if len(errorMap) > 0 {
		return dto.UserResponse{}, &AppErrorpkg.AppError{
			Message: "Validation error",
			Code:    "VALIDATION_ERROR",
			Status:  422,
			Errors:  errorMap,
		}
	}
	hashedPassword, err := s.hasher.Hash(req.Password)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("hashed password failed: %w", err)
	}
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := s.userRepo.Create(&user); err != nil {
		return dto.UserResponse{}, fmt.Errorf("create user failed: %w", err)
	}
	return mappers.ToUserResponseDetail(user), nil
}

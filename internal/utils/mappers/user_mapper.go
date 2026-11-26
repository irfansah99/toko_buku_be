package mappers

import (
	"backend-toko-buku/internal/dto"
	"backend-toko-buku/internal/models"
)

func ToUserResponseDetail(u models.User) dto.UserResponse {
	return dto.UserResponse{
		Id:        u.Id,
		Name:      u.Name,
		Image:     u.Image,
		Role:      u.Role,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

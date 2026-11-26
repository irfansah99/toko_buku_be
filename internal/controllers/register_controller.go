package controllers

import (
	"backend-toko-buku/internal/dto"
	"backend-toko-buku/internal/services/auth"
	AppErrorpkg "backend-toko-buku/internal/services/errors"
	"backend-toko-buku/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	service *auth.RegisterService
}

func NewRegisterController(service *auth.RegisterService) *RegisterController {
	return &RegisterController{service}
}
func (ctrl *RegisterController) Register(c *gin.Context) {

	var req dto.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  utils.TranslateErrorMessage(err),
		})
		return
	}

	user, err := ctrl.service.RegisterUser(req)

	if err != nil {

		if appErr, ok := err.(*AppErrorpkg.AppError); ok {
			c.JSON(appErr.Status, dto.ErrorResponse{
				Success: false,
				Message: appErr.Message,
				Code:    appErr.Code,
				Errors:  appErr.Errors,
			})
			return
		}

		log.Println("ERROR Register:", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to update user",
			Code:    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

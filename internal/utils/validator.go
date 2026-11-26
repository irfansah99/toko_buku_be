package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10" // Mengimpor package untuk validasi
	// Mengimpor GORM untuk menangani error terkait database
)

// TranslateErrorMessage menangani validasi dari validator.v10 dan duplikasi entri dari GORM
func TranslateErrorMessage(err error) map[string]string {
	// Membuat map untuk menampung pesan error
	errorsMap := make(map[string]string)

	// Handle validasi dari validator.v10
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field() // Menyimpan nama field yang gagal validasi
			switch fieldError.Tag() {   // Menangani berbagai jenis validasi
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field) // Pesan error jika field kosong
			case "email":
				errorsMap[field] = "Invalid email format" // Pesan error jika format email tidak valid
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param()) // Pesan error jika nilai terlalu pendek
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param()) // Pesan error jika nilai terlalu panjang
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field) // Pesan error jika nilai bukan angka
			default:
				errorsMap[field] = "Invalid value" // Pesan error default untuk kesalahan validasi lainnya
			}
		}
	}

	// Mengembalikan map yang berisi pesan error
	return errorsMap
}

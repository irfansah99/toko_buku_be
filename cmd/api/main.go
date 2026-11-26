package main

import (
	"backend-toko-buku/config"
	"backend-toko-buku/database"

	"backend-toko-buku/internal/controllers"
	"backend-toko-buku/internal/models"
	"backend-toko-buku/internal/repository"
	"backend-toko-buku/internal/routes"
	"backend-toko-buku/internal/services/auth"
	"backend-toko-buku/internal/services/hash"
)

func main() {


	config.LoadEnv()

	db := database.InitDB()


	db.AutoMigrate(&models.User{})
	hasher := &hash.BcryptHasher{}
	userRepo := repository.NewUserRepository(db)

	registerService := auth.NewRegisterService(userRepo, hasher)

	registerController := controllers.NewRegisterController(registerService)

	r := routes.SetupRouter(registerController)

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}

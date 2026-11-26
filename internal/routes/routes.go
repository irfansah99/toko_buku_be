package routes

import (
	"backend-toko-buku/internal/controllers"
	"backend-toko-buku/internal/middelwares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	registerCtrl *controllers.RegisterController,

) *gin.Engine {
	router := gin.Default()

	router.Static("/uploads", "./uploads")

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")

	// public
	AuthRoutes(api, registerCtrl)

	// protected
	auth := api.Group("/")
	auth.Use(middelwares.AuthMiddleware())

	return router
}

package main

import (
	"backend-toko-buku/config"
	"backend-toko-buku/database"

	"backend-toko-buku/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {

	//load config .env
	config.LoadEnv()

	//inisialiasai Gin
	router := gin.Default()
	db := database.InitDB()
	// rdb := config.RedisClient

	// cache := cache.NewRedisCache(rdb)
	db.AutoMigrate(&models.User{})
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//mulai server
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}

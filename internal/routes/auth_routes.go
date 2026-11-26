package routes

import (
	"backend-toko-buku/internal/controllers"


	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, registerctrl *controllers.RegisterController) {
	r.POST("/register", registerctrl.Register)


}

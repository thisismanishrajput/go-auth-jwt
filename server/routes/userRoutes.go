package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/thisismanishrajput/goauth/server/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.Use()
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}

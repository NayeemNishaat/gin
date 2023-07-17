package route

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(ar *gin.RouterGroup) {
	router := ar.Group("/auth")

	router.GET("/basic", middleware.Auth(), controller.BasicAuth)

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}

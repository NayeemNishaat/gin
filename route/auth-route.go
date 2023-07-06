package route

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(gr *gin.RouterGroup) {
	router := gr.Group("/auth")

	router.GET("/basic", middleware.Auth(), controller.BasicAuth)

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}

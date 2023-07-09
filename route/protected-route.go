package route

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func User(gr *gin.RouterGroup) {
	router := gr.Group("/user")
	router.Use(middleware.JwtAuthMiddleware())

	router.GET("/", controller.Me)
}

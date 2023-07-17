package route

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func User(ar *gin.RouterGroup) {
	router := ar.Group("/user")
	router.Use(middleware.JwtAuthMiddleware())

	router.GET("/", controller.Me)
}

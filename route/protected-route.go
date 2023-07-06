package route

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func Protected(gr *gin.RouterGroup) {
	router := gr.Group("/protected")
	router.Use(middleware.JwtAuthMiddleware())
	router.GET("/", controller.BasicAuth)
}

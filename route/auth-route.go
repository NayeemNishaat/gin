package route

import (
	"gin/controller"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(s *gin.RouterGroup) {
	router := s.Group("/auth")

	router.POST("/basic", middleware.Auth(), controller.CreateAlbum)
}

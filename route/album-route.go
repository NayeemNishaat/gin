package route

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(s *gin.Engine) {
	router := s.Group("/album")

	router.GET("/", controller.GetAlbums)
	router.GET("/:id", controller.GetAlbumByID)
	router.POST("/", controller.PostAlbums)
}

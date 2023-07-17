package route

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(ar *gin.RouterGroup) {
	router := ar.Group("/album")

	router.GET("/", controller.GetAllAlbums)
	router.GET("/:id", controller.GetAlbumByID)
	router.POST("/", controller.CreateAlbum)
}

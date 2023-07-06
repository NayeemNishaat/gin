package route

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(gr *gin.RouterGroup) {
	router := gr.Group("/album")

	router.GET("/", controller.GetAllAlbums)
	router.GET("/:id", controller.GetAlbumByID)
	router.POST("/", controller.CreateAlbum)
}

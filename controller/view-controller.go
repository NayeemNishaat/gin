package controller

import (
	"gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	albums := service.GetAllAlbums()
	data := gin.H{
		"title":  "Home Page",
		"albums": albums,
	}

	c.HTML(http.StatusOK, "index.html", data)
}

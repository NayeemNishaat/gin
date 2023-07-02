package main

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/albums", controller.PostAlbums)

	router.Run("localhost:3000")
	// router.Run(":3000")
	// router.Run()
}

// nodemon -q -e go --signal SIGTERM --exec go run .

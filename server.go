package main

import (
	"gin/route"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	route.AlbumRoutes(server)

	server.Run("localhost:3000")
	// router.Run(":3000")
	// router.Run()
}

// nodemon -q -e go --signal SIGTERM --exec go run .

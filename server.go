package main

import (
	"gin/lib"
	"gin/model"
	"gin/route"

	"github.com/gin-gonic/gin"
)

func main() {
	lib.StoreLog()
	model.ConnectDataBase()

	server := gin.Default()

	lib.MountFuncMap(server)
	server.Static("/style", "./public/style")
	server.LoadHTMLGlob("./public/template/*.html")

	apiRouter := server.Group("/api/v1")
	viewRouter := server.Group("/")

	// Chapter: APIs
	route.AlbumRoutes(apiRouter)
	route.AuthRoutes(apiRouter)
	route.User(apiRouter)

	// Chapter: Views
	route.View(viewRouter)

	server.Run("localhost:3000")
	// router.Run(":3000")
	// router.Run()
}

// nodemon -q -e go --signal SIGTERM --exec go run .

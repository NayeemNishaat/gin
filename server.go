package main

import (
	"fmt"
	"gin/model"
	"gin/route"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func storeLog() {
	// f, err := os.OpenFile("./log/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Note: Append
	f, err := os.Create("./log/server.log") // Note: Create New and Append

	if err != nil {
		fmt.Println("Open Log File Failed", err)
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.ForceConsoleColor()
}

func main() {
	storeLog()
	model.ConnectDataBase()

	server := gin.Default()

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

package main

import (
	"fmt"
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

	server := gin.Default()
	globalRouter := server.Group("/api/v1")

	route.AlbumRoutes(globalRouter)
	route.AuthRoutes(globalRouter)

	server.Run("localhost:3000")
	// router.Run(":3000")
	// router.Run()
}

// nodemon -q -e go --signal SIGTERM --exec go run .

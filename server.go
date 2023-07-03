package main

import (
	"fmt"
	"gin/route"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func storeLog() {
	f, err := os.Create("./log/server.log")

	if err != nil {
		fmt.Println("Open Log File Failed", err)
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.ForceConsoleColor()
}

func main() {
	storeLog()

	server := gin.Default()

	route.AlbumRoutes(server)

	server.Run("localhost:3000")
	// router.Run(":3000")
	// router.Run()
}

// nodemon -q -e go --signal SIGTERM --exec go run .

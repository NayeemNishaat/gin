package lib

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func StoreLog() {
	// f, err := os.OpenFile("./log/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Note: Append
	f, err := os.Create("./log/server.log") // Note: Create New and Append

	if err != nil {
		fmt.Println("Open Log File Failed", err)
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.ForceConsoleColor()
}

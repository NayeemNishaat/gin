package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Success",
	})
}

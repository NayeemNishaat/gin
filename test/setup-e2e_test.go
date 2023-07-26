package test

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	router := gin.Default()
	return router
}

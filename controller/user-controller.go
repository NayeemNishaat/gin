package controller

import (
	"gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	// userId, err := lib.ExtractTokenID(c)
	uId, _ := c.Get("userId")

	if uId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "user doesn't exist"})
		return
	}

	// fmt.Println(reflect.TypeOf(uId)) // Note: Check the type of a var

	userId, ok := uId.(uint64)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "invalid user id"})
		return
	}

	u, err := service.Me(uint(userId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": true, "message": "success", "data": u})
}

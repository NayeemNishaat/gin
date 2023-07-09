package controller

import (
	"gin/lib"
	"gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Success",
	})
}

func Me(c *gin.Context) {
	// userId, err := lib.ExtractTokenID(c)
	uId, _ := c.Get("userId")

	if uId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "user doesn't exist"})
		return
	}

	userId, ok := uId.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "invalid user id"})
		return
	}

	u, err := service.Me(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": true, "message": "success", "data": u})
}

type RegisterData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var rd RegisterData

	errCustomizer := lib.GetCustomizer(rd)

	if err := c.ShouldBindJSON(&rd); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Request body is missing!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error(), "validation": errCustomizer.DecryptErrors(err)})
		}
		return
	}

	_, err := service.Register(rd.Username, rd.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": true, "message": "registration success"})
}

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var ld LoginData

	errCustomizer := lib.GetCustomizer(ld)

	if err := c.ShouldBindJSON(&ld); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Request body is missing!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error(), "validation": errCustomizer.DecryptErrors(err)})
		}
		return
	}

	token, err := service.Login(ld.Username, ld.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success!", "data": map[string]string{"token": token}})
}

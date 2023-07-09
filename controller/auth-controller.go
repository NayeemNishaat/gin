package controller

import (
	"gin/lib"
	"gin/model"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "user doesn't exist"})
		return
	}

	userId, ok := uId.(uint)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	u, err := model.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is missing!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error(), "validation": errCustomizer.DecryptErrors(err)})
		}
		return
	}

	u := model.User{}

	u.Username = rd.Username
	u.Password = rd.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var ld LoginData

	if err := c.ShouldBindJSON(&ld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{}

	u.Username = ld.Username
	u.Password = ld.Password

	token, err := model.ValidateLogin(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

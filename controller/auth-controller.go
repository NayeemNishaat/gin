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
	user_id, err := lib.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := model.GetUserByID(user_id)

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

	if err := c.ShouldBindJSON(&rd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

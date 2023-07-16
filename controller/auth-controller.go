package controller

import (
	"gin/lib"
	"gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

func BasicAuth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Success",
	})
}

type RegisterData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" g:"isStrong" isStrong:"$field should contain at least a special character, a number, a uppercase letter and minimum 8 characters long" binding:"required" required:"$field is required"`
	// Password string `json:"password" validate:"isStrong" binding:"required" required:"$field is required"`
	// Email string `json:"email" xml:"email" form:"email" binding:"required,min=3,max=10,email"`
	// URL string `json:"url" binding:"required,url"`
	// Age int8 `json:"age" binding:"gte=10,lte=100"`
}

func Register(c *gin.Context) {
	// validate := validator.New()
	// validate.RegisterValidation("isStrong", lib.ValidateStrongPass)

	var rd RegisterData

	CustomValidators := galidator.Validators{"isStrong": lib.ValidateStrongPass}

	errCustomizer := lib.GetCustomCustomizer(rd, CustomValidators)

	if err := c.ShouldBindJSON(&rd); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Request body is missing!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error(), "validation": errCustomizer.DecryptErrors(err)})
		}
		return
	}

	// error := validate.Struct(rd)
	error := errCustomizer.Validate(rd)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Validation Failed", "validation": error})
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

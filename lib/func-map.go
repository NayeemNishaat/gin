package lib

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func MountFuncMap(s *gin.Engine) {
	s.SetFuncMap(template.FuncMap{
		"sub": sub,
	})
}

func sub(a int, b int) int {
	return a - b
}

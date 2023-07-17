package route

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func View(vr *gin.RouterGroup) {
	vr.GET("/", controller.Index)
}

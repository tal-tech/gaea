package router

import (
	"gaea/app/controller/democontroller"
	"github.com/gin-gonic/gin"
)

//The routing method is exactly the same as Gin
func RegisterRouter(router *gin.Engine) {
	entry := router.Group("/demo")
	entry.GET("/test", democontroller.GaeaDemo)
}

package routes

import (
	"gowebdemo/internal/app/appone/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckAuth = append(routerNoCheckAuth, registerDebugRouter)
}

func registerDebugRouter(v1 *gin.RouterGroup) {
	c := controllers.Debug{}

	v1.GET("/debug", c.Debug)
}

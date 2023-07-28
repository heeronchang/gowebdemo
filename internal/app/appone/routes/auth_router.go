package routes

import (
	"gowebdemo/internal/app/appone/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckAuth = append(routerNoCheckAuth, registerAuthRouter)
}

func registerAuthRouter(v1 *gin.RouterGroup) {
	c := &controllers.Auth{}

	v1.POST("/login", c.Login)
}

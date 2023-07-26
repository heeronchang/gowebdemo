package routes

import (
	"context"
	"gowebdemo/configs/appone"
	"gowebdemo/internal/pkg/middleware"
	"gowebdemo/internal/pkg/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 启动普通webapi
func StartWebAPI() {
	addr := appone.GetAddr()
	r := routers()

	srv := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	var shutdown chan bool = make(chan bool)
	utils.HandleShutdown(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Server Shutdown gracefully err: %s\n", err.Error())
		}
		log.Printf("Server shutdown.")
		shutdown <- true
	})

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen addr: %s err: %s\n", addr, err.Error())
	}

	<-shutdown
}

type AuthMiddleWare func() gin.HandlerFunc

var (
	routerCheckAuth   = make([]func(*gin.RouterGroup, AuthMiddleWare), 0)
	routerNoCheckAuth = make([]func(*gin.RouterGroup), 0)
)

func routers() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MB

	// 中间件
	r.Use(middleware.Cors(), middleware.JWTAuth())
	r.Use(middleware.RequestID(), middleware.ZeroLog())

	// routers
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong!")
	})

	checkAuthRouter(r, middleware.JWTAuth)
	noCheckAuthRouter(r)

	return r
}

func checkAuthRouter(r *gin.Engine, authMiddleWare AuthMiddleWare) {
	v1 := r.Group("/api/v1")

	for _, f := range routerCheckAuth {
		f(v1, authMiddleWare)
	}
}

func noCheckAuthRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	for _, f := range routerNoCheckAuth {
		f(v1)
	}
}

func StartWebsocketAPI() {

}

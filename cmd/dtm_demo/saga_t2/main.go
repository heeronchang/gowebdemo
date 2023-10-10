package main

import (
	"context"
	"gowebdemo/internal/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	// routers
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong!")
	})

	in := r.Group("/api/bank2")

	in.POST("/TransIn", func(c *gin.Context) {
		log.Printf("TransIn")
		c.JSON(200, "")
		// c.JSON(409, "") // Status 409 for Failure. Won't be retried
	})
	in.POST("/TransInCompensate", func(c *gin.Context) {
		log.Printf("TransInCompensate")
		c.JSON(200, "")
	})

	addr := "0.0.0.0:8083"
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
			log.Info().Msgf("Server Shutdown gracefully err: %s\n", err.Error())
		}
		log.Info().Msgf("Server shutdown.")
		shutdown <- true
	})

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Info().Msgf("listen addr: %s err: %s\n", addr, err.Error())
	}

	<-shutdown
}

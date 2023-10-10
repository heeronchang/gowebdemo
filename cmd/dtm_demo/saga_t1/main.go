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
	r.MaxMultipartMemory = 8 << 20 // 8 MB

	// routers
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong!")
	})

	out := r.Group("/api/bank1")

	out.POST("/TransOut", func(c *gin.Context) {
		log.Printf("TransOut")
		c.JSON(200, "")
	})
	out.POST("/TransOutCompensate", func(c *gin.Context) {
		log.Warn().Msg("TransOutCompensate")
		// log.Printf("TransOutCompensate")
		c.JSON(200, "")
	})

	addr := "0.0.0.0:8082"

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

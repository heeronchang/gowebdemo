package controllers

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Debug struct{}

func (d *Debug) Debug(ctx *gin.Context) {
	requestId := requestid.Get(ctx)
	log.Info().Msgf("request id:%s", requestId)
	ctx.JSON(http.StatusOK, gin.H{"message": "debug success", "id": requestId})
}

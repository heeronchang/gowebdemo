package controllers

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type Debug struct{}

func (d *Debug) Debug(ctx *gin.Context) {
	requestId := requestid.Get(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "debug success", "id": requestId})
}

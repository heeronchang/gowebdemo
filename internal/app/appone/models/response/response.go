package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data interface{}, msg any) {
	if msg == nil {
		msg = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": msg,
		"data":    data,
	})
}

func Fail(c *gin.Context, msg any) {
	if msg == nil {
		msg = "error"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
	})
}

package middleware

import (
	"gowebdemo/configs/appone"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Cors() gin.HandlerFunc {
	corsCfg := cors.DefaultConfig()

	corsCfg.AllowAllOrigins = true
	corsCfg.AllowCredentials = true
	corsCfg.AddAllowHeaders("token", "refresh_token")

	return cors.New(corsCfg)
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var h map[string]string
		if err := ctx.BindHeader(&h); err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		}
		// TODO 验证token

		// 如果中间件里读取了request body，需要把body数据复制一份，重新放入ctx
		// data, err := ctx.GetRawData()
		// if err != nil {
		// 	ctx.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }
		// ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data)) // 复制一份数据塞回 ctx，避免handler读取不到数据

		// ctx.Next()
	}
}

func ZeroLog() gin.HandlerFunc {
	logConf := appone.GetLogger()
	file, err := os.OpenFile(logConf.OUTPUT, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Msgf("创建或打开日志文件(%s)失败：%s", logConf.OUTPUT, err.Error())
	}

	return logger.SetLogger(
		logger.WithLogger(func(ctx *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(file).With().Str("id", requestid.Get(ctx)).Logger()
		}),
	)
}

func RequestID() gin.HandlerFunc {
	return requestid.New(requestid.WithGenerator(func() string {
		uId := uuid.New().String()
		return uId
	}), requestid.WithCustomHeaderStrKey("X-Request-ID"))
}

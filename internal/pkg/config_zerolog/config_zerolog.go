package configzerologger

import (
	"gowebdemo/configs/appone"
	"gowebdemo/internal/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	logConf := appone.GetLogger()
	filePath := logConf.OUTPUT
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		log.Fatalf("创建日志外层文件夹(%s)失败：%s", filepath.Dir(filePath), err.Error())
	}

	file, err := os.OpenFile(logConf.OUTPUT, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("创建或打开日志文件(%s)失败：%s", logConf.OUTPUT, err.Error())
	}
	// defer file.Close()
	utils.HandleShutdown(func() {
		err = file.Close()
		if err != nil {
			log.Printf("zero logger file closed err:%s", err.Error())
		} else {
			log.Println("zero logger file closed.")
		}
	})

	// 日志归档
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    logConf.MaxSize,
		MaxBackups: logConf.MaxBackups,
		MaxAge:     logConf.MaxAge,
		Compress:   logConf.Compress,
	}

	level := string2Level(logConf.LEVEL)
	var multiWriter zerolog.LevelWriter
	if utils.IsDirectlyRun() {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		multiWriter = zerolog.MultiLevelWriter(consoleWriter, lumberjackLogger)
	} else {
		multiWriter = zerolog.MultiLevelWriter(lumberjackLogger)
	}

	zerolog.SetGlobalLevel(level)
	// zerolog.DefaultContextLogger.With().Stack().Caller().Timestamp().CallerWithSkipFrameCount(2)
	zlog.Logger = zlog.Output(multiWriter).With().CallerWithSkipFrameCount(2).Logger()
}

// 字符串转日志等级
func string2Level(level string) zerolog.Level {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARNING":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "PANIC":
		return zerolog.PanicLevel
	case "FATAL":
		return zerolog.FatalLevel
	default:
		return zerolog.Disabled
	}
}

package logger

import (
	"gowebdemo/configs/appone"
	"gowebdemo/internal/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	PANIC
	FATAL
	UNKNOWN
)

var SingleLogger *log.Logger
var logLevel = INFO

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
			log.Printf("logger file closed err:%s", err.Error())
		} else {
			log.Println("logger file closed.")
		}
	})

	SingleLogger = log.New(os.Stderr, "", log.LstdFlags)
	SingleLogger.SetOutput(file)

	// 设置日志等级
	level := string2Level(logConf.LEVEL)
	SetLogLevel(level)
}

// 设置日志等级
func SetLogLevel(level LogLevel) {
	logLevel = level
}

// 字符串转日志等级
func string2Level(level string) LogLevel {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	case "PANIC":
		return PANIC
	case "FATAL":
		return FATAL
	default:
		return UNKNOWN
	}
}

func Debug(v ...any) {
	if logLevel > DEBUG {
		return
	}
	SingleLogger.SetPrefix("[DEBUG]")
	SingleLogger.Print(v...)
}

func Debugf(format string, v ...any) {
	if logLevel > DEBUG {
		return
	}
	SingleLogger.SetPrefix("[DEBUG]")
	SingleLogger.Printf(format, v...)
}

func Info(v ...any) {
	if logLevel > INFO {
		return
	}
	SingleLogger.SetPrefix("[INFO]")
	SingleLogger.Println(v...)
}

func Infof(format string, v ...any) {
	if logLevel > INFO {
		return
	}
	SingleLogger.SetPrefix("[INFO]")
	SingleLogger.Printf(format, v...)
}

func Warn(v ...any) {
	if logLevel > WARNING {
		return
	}
	SingleLogger.SetPrefix("[WARNING]")
	SingleLogger.Println(v...)
}

func Warnf(format string, v ...any) {
	if logLevel > WARNING {
		return
	}
	SingleLogger.SetPrefix("[WARNING]")
	SingleLogger.Printf(format, v...)
}

func Error(v ...any) {
	if logLevel > ERROR {
		return
	}
	SingleLogger.SetPrefix("[ERROR]")
	SingleLogger.Println(v...)
}

func Errorf(format string, v ...any) {
	if logLevel > ERROR {
		return
	}
	SingleLogger.SetPrefix("[ERROR]")
	SingleLogger.Printf(format, v...)
}

func Panic(v ...any) {
	if logLevel > PANIC {
		return
	}
	SingleLogger.SetPrefix("[PANIC]")
	SingleLogger.Panic(v...)
}

func Panicf(format string, v ...any) {
	if logLevel > PANIC {
		return
	}
	SingleLogger.SetPrefix("[PANIC]")
	SingleLogger.Panicf(format, v...)
}

func Fatal(v ...any) {
	if logLevel > FATAL {
		return
	}
	SingleLogger.SetPrefix("[FATAL]")
	SingleLogger.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	if logLevel > FATAL {
		return
	}
	SingleLogger.SetPrefix("[FATAL]")
	SingleLogger.Fatalf(format, v...)
}

package appone

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var AppOneConfig *viper.Viper

func init() {
	AppOneConfig = viper.New()

	runConfPath := "configs/appone/"
	binConfPath := "configs/"

	var confPath string

	executableName := os.Args[0]
	log.Printf("executableName:%s\n", executableName)

	// 判断程序是直接运行还是作为二进制包运行
	if ok := strings.HasPrefix(executableName, "./"); ok {
		log.Println("作为二进制包运行")
		confPath = binConfPath
	} else {
		log.Println("直接运行")
		confPath = runConfPath
	}

	AppOneConfig.AddConfigPath(runConfPath)
	AppOneConfig.AddConfigPath(binConfPath)
	AppOneConfig.SetConfigType("yaml")

	if err := AppOneConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// "fatal error config file(configs/appone/config.yaml) not found"
			panic(fmt.Errorf("fatal error config file(%s) not found", confPath))
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	AppOneConfig.OnConfigChange(func(in fsnotify.Event) {
		log.Println("Config file changed:", in.Name)
	})
	AppOneConfig.WatchConfig()

}

// GetTitle 获取标题
func GetTitle() string {
	return AppOneConfig.GetString("title")
}

// GetSubTitle 获取副标题
func GetSubTitle() string {
	return AppOneConfig.GetString("sub_title")
}

// GetDesc 获取描述
func GetDesc() string {
	return AppOneConfig.GetString("description")
}

// GetIP 获取程序监听IP
func GetIP() string {
	return AppOneConfig.GetString("net.ip")
}

// GetPort 获取程序监听端口
func GetPort() int64 {
	return AppOneConfig.GetInt64("net.port")
}

// GetWebsocketPort 获取websocket 端口
func GetWebsocketPort() int64 {
	return AppOneConfig.GetInt64("net.websocket")
}

// GetAddr 获取程序监听地址
func GetAddr() string {
	return fmt.Sprintf("%s:%d", GetIP(), GetPort())
}

// GetTimeout 获取网络超时时间
func GetTimeout() int64 {
	return AppOneConfig.GetInt64("net.timeout")
}

// LoggerConfig 日志配置结构体
type LoggerConfig struct {
	OUTPUT     string // 日志输出路径
	LEVEL      string // 日志等级
	MaxSize    int    // 单个日志文件 大小上限，单位：MB
	MaxBackups int    // 最多保留的旧日志个数
	MaxAge     int    // 保留日志最大天数
	Compress   bool   //  是否启用压缩
}

// GetLogger 获取日志配置
func GetLogger() *LoggerConfig {
	lc := &LoggerConfig{
		OUTPUT:     AppOneConfig.GetString("logger.OUTPUT"),
		LEVEL:      AppOneConfig.GetString("logger.LEVEL"),
		MaxSize:    AppOneConfig.GetInt("logger.MAX_SIZE"),
		MaxBackups: AppOneConfig.GetInt("logger.MAX_BACKUPS"),
		MaxAge:     AppOneConfig.GetInt("logger.MAX_AGE"),
		Compress:   AppOneConfig.GetBool("logger.COMPRESS"),
	}

	return lc
}

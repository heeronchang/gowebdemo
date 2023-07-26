package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// 监听程序中断信号，清理并退出
type Callback func()

var cbs []Callback
var sigChan chan os.Signal

func init() {
	cbs = make([]Callback, 0)
	sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Printf("接收到信号：%v", sig)

		for _, cb := range cbs {
			cb()
		}
		log.Printf("already call all cbs")
	}()
}

// HandleShutdown 监听信号执行清理回调函数
func HandleShutdown(cb Callback) {
	cbs = append(cbs, cb)
}

package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// IsDirectlyRun 程序是否是直接运行
func IsDirectlyRun() bool {
	executableName := os.Args[0]
	log.Printf("executableName:%s\n", executableName)

	// 判断程序是直接运行还是作为二进制包运行
	if ok := strings.HasPrefix(executableName, "./"); ok {
		log.Println("作为二进制包运行")
		return true
	} else {
		fmt.Println("直接运行")
		return false
	}
}

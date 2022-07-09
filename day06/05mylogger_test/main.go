package main

import (
	"fmt"
	"time"

	"github.com/wanghengg/mylogger"
)

func main() {
	// log := mylogger.NewLog("Error")
	log, err := mylogger.NewFileLogger("warning", "./", "20220709.log", 10*1024*1024)
	if err != nil {
		fmt.Printf("NewFileLogger failed:%v", err)
		return
	}
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}

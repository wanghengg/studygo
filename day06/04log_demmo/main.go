package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("这是一个测试的日志")
		time.Sleep(2 * time.Second)
	}
}

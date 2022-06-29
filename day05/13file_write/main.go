package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开文件读写

func writeDemo() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	// write
	fileObj.Write([]byte("zhoulin menngbi le!\n"))
	fileObj.WriteString("周林解释不了\n")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer fileObj.Close()

	// 创建一个读写对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello world!\n") // 写到缓存中
	wr.Flush()                       // 将缓存中的内容刷新到文件
}

func main() {
	writeDemo()
	writeDemo2()
}

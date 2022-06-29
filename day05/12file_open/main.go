package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 打开文件

func readFromFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()

	// 读文件
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Printf("读了%d个字节\n", n)
			fmt.Println(string(tmp[:n]))
			break
		}
		if err != nil {
			fmt.Printf("read from file failed: err:%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

// 利用bufio这个包读取文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()

	// 创建一个用于从文件读取内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read from file failed: %v\n", err)
		}
		fmt.Print(line)
	}
}

func main() {
	readFromFile()

	fmt.Println("............................")

	readFromFilebyBufio()
}

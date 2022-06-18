package main

import (
	"fmt"
	"os"
	"strings"
)

// 使用递归函数打印指定文件夹下的所有文件
func ListFiles(dir string) {
	dirList, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, val := range dirList {
		if strings.HasPrefix(val.Name(), ".") {
			continue
		}
		fmt.Println(val.Name())
		if val.IsDir() {
			ListFiles(strings.Join([]string{dir, val.Name()}, "/"))
		}
	}
}

func main() {
	dirname := "../../"
	ListFiles(dirname)
}

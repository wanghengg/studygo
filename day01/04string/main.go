package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/home/wangheng/workspace/go/studygo/day01"
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	// 反单引号保持字符串不变
	s3 := `/home/wangheng/workspace/go/studygo/day01`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	hello := "hello"
	world := "world"
	ss := hello + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", hello, world)
	fmt.Println(ss1)

	// 字符串分割
	ret := strings.Split(path, "/")
	fmt.Println(ret)
}

package main

import "fmt"

var x = 100

func f1() {
	fmt.Println(x)
}

func f2() {
	// 函数运行时先从局部变量中查找变量，如果找不到再从全局变量中找（从内部到外部）
	x := 10
	fmt.Println(x)
}

func main() {
	f1()
	f2()
}

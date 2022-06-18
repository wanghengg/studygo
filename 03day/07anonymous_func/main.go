package main

import "fmt"

// 匿名函数

func main() {
	// 函数内部无法声明带名字的函数
	// 匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(100, 200)
	// 如果只是调用一次的函数，可以写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("Hello world!")
	}(100, 200)
}

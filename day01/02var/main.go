package main

import "fmt"

// // 声明变量
// var name string
// var age int
// var isOK bool

// 批量声明，推荐使用驼峰声明
var (
	name string
	age  int
	isOK bool
)

// l := 5

func main() {
	name = "理想"
	age = 35
	isOK = true
	// go语言中非全局变量声明了必须使用，不使用无法通过编译
	fmt.Println(name, age, isOK)

	// 简短变量声明只能用于函数内，不能用于全局变量声明
	n := 10
	fmt.Println(n)
}

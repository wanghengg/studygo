package main

import "fmt"

func f1() {
	fmt.Println("hello world！")
}

func f2(name string) {
	fmt.Println("Hello", name)
}

func f3(x, y int) int {
	sum := x + y
	return sum
}

// 可变参数
func f5(title string, y ...int) {
	fmt.Println(title)
	fmt.Println(y)
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 使用命名返回值，可以在函数中直接使用该变量，不需要声明
	return      // 已经指定返回值变量，不需要写出返回值变量
}

func main() {
	f1()
	f2("xiaowang")
	fmt.Println(f3(2, 3))
	f5("xiaowang", 1, 2, 3, 4)
	fmt.Println(f6(2, 3))
}

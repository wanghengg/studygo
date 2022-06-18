package main

import "fmt"

//defer
// Go语言中的return语句不是原子操作，分为两步：
// 第一步：返回值赋值
// 第二部：返回结果
// 存在defer语句时，将defer语句插在两步之间

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func main() {
	fmt.Println("start")
	defer fmt.Println("哈哈哈") // defer关键字修饰的语句函数退出的时候执行
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("嘎嘎嘎") // 多个defer语句，遵守先进后出的规则
	fmt.Println("end")

	fmt.Println(f1())
	fmt.Println(f2())
}

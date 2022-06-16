package main

import "fmt"

// 常量
// 定义之后不能修改
const pi = 3.1415926

// 批量声明
const (
	statusOk = 200
	notFound = 404
)

// 批量声明常量时，如果某行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota
const (
	a1 = iota // iota出现时为0
	a2
	a3
)

func main() {
	fmt.Println(pi)
	fmt.Println(statusOk, notFound)
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3)
	fmt.Printf("%T\n", pi)
}

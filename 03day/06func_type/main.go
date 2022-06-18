package main

import (
	"fmt"
)

func f1() {
	fmt.Println("Hello, world!")
}

func f2() int {
	return 10
}

// 函数作为参数
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f2)
	fmt.Println(f5(f2)(2, 3))
}

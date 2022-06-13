package main

import "fmt"

func main() {
	// 1. &:取地址
	// 2. *:根据地址取值
	n := 128
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	var a *int
	fmt.Println(a)
	a = new(int)
	*a = 100
	fmt.Println(*a)

	x := make(map[int]int, 2)
	fmt.Println(x)
}

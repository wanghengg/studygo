package main

import "fmt"

// 自定义类型和类型别名

type myInt int    // 自定义类型
type youInt = int // 类型别名

func main() {
	var n myInt = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m youInt = 200
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}

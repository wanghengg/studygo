package main

import (
	"fmt"
)

func main() {
	s := "Hello沙河"
	n := len(s)
	fmt.Println(n)

	// for _, c := range s {
	// 	fmt.Printf("%c\n", c)
	// }

	// 字符串修改
	s2 := "白萝卜"
	fmt.Println(s2)
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	c3 := 'a'
	fmt.Printf("c1:%T c2:%T c3:%T\n", c1, c2, c3)
}

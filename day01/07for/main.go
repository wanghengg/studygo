package main

import "fmt"

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for range循环
	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d, %c\n", i, v)
	}
}

package main

import "fmt"

// 空接口

func show(v interface{}) {
	fmt.Printf("type: %T, value: %v\n", v, v)
}

func main() {
	m1 := make(map[string]interface{}, 16)

	m1["name"] = "周林"
	m1["age"] = 25
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}

	fmt.Println(m1)

	show(m1)
}

package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var xiaowang Person = Person{name: "xiaowang", age: 25, gender: "male", hobby: []string{"basketball", "ping-pong"}}
	fmt.Println(xiaowang)

	// 匿名结构体，临时使用
	var s struct {
		name string
		age  int
	}
	s.name = "xiaoli"
	s.age = 12
	fmt.Printf("type:%T value:%v\n", s, s)
}

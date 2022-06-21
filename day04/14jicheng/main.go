package main

import (
	"fmt"
)

// 用结构体模拟继承

type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会移动\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

// 给dog实现一个汪汪汪的method
func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪\n", d.name)
}

func main() {
	d := dog{
		animal: animal{
			name: "wang",
		},
		feet: 12,
	}
	d.wang()
}

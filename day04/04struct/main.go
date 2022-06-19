package main

import "fmt"

type person struct {
	name, gengder string
}

func f(x person) {
	x.gengder = "女" // go语言函数传参是拷贝
}

func f2(x *person) {
	(*x).gengder = "女"
}

func main() {
	var p person
	p.name = "xiaowang"
	p.gengder = "男"
	f(p)
	fmt.Println(p.gengder)
	f2(&p)
	fmt.Println(p.gengder)

	// 结构体指针1
	var p2 = new(person)
	p2.name = "yuan"
	fmt.Printf("%T\n", p2)
	fmt.Println(p2)
	fmt.Printf("%p\n", p2)

	// 结构体指针2
	var p3 = &person{
		name:    "元帅",
		gengder: "男",
	}
	fmt.Printf("%#v\n", p3)
}

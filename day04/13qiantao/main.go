package main

import (
	"fmt"
)

// 嵌套结构体

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套
	work    workPlace
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "周林",
		age:  90000,
		address: address{
			province: "湖北",
			city:     "武汉",
		},
		work: workPlace{},
	}
	fmt.Println(p1)
	fmt.Println(p1.city)
	fmt.Println(p1.address.province)

	com := company{
		name: "腾讯",
		addr: address{
			province: "广东",
			city:     "深圳",
		},
	}

	fmt.Println(com)
}

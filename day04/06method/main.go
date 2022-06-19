package main

import (
	"fmt"
)

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
}

// 值接收者：拷贝
func (p person) printAge() {
	fmt.Println(p.age)
}

// 指针接收者：传地址
func (p *person) guonian() {
	p.age++
}

func main() {
	d1 := newDog("zhoulin")
	d1.wang()

	p1 := newPerson("元帅", 18)
	fmt.Println(p1.age)
	p1.printAge()
	p1.guonian()
	fmt.Println(p1.age)
	p1.printAge()
}

package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步...")
}

// cat实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	var e1 eater
	c1 := cat{"miao", 4}
	e1 = &c1

	var m1 mover = &c1

	e1.eat("fish")
	m1.move()

	var a1 animal = &c1
	a1.eat("fish")
	a1.move()
}

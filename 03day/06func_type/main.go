package main

import "fmt"

func f1() {
	fmt.Println("Hello, world!")
}

func f2() int {
	return 10
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
}

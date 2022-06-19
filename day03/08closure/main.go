package main

import (
	"fmt"
)

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo val = %d\n", *x)
	}
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %d\n", val)
		}()
	}
}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func main() {
	x := 133
	f1 := foo1(&x)
	f2 := foo2(x)
	f1()
	f2()
	f1()
	f2()

	x = 233
	f1()
	f2()
	f1()
	f2()

	foo1(&x)()
	foo2(x)()
	foo1(&x)()
	foo2(x)()

	fmt.Println("foo3()---------------------")
	foo3()

	fmt.Println("foo4()----------------------")
	foo4()

	fmt.Println("foo5()-----------------------")
	foo5()

	fmt.Println("foo7()----------------------")
	f7s := foo7(11)
	for _, f7 := range f7s {
		f7()
	}
}

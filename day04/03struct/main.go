package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var xiaowang Person = Person{name: "xiaowang", age: 25, gender: "male", hobby: []string{"basketball", "ping-pong"}}
	fmt.Println(xiaowang)
}

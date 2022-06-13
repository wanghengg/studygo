package main

import "fmt"

func main() {
	var m1 map[string]int

	fmt.Println(m1 == nil)

	m1 = make(map[string]int, 10)

	m1["理想"] = 18
	m1["小王"] = 34
	fmt.Println(m1)

	fmt.Println(m1["理想"])

	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	delete(m1, "理想")
	fmt.Println(m1)
}

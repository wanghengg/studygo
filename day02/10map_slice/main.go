package main

import "fmt"

func main() {
	var s1 = make([]map[int]string, 1, 10)

	s1[0] = make(map[int]string, 1)

	s1[0][10] = "A"
	fmt.Println(s1)
}

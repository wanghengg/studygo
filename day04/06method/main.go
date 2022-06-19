package main

import "fmt"

type dog struct {
	name string
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

func main() {
	d1 := newDog("zhoulin")
	d1.wang()
}

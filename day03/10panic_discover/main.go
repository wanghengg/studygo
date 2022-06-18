package main

import "fmt"

// panic和recover

func funcA() {
	fmt.Println("a")
}

func funcB() {
	// 刚刚打开的连接
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重错误")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}

package main

import "fmt"

// if条件判断
func main() {
	age := 39
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("未成年！")
	}

	// 多个条件判断
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

	// age1是局部变量，if判断结束释放内存
	if age1 := 17; age1 > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}

package main

import (
	"fmt"
)

// 类型断言

func assign(v interface{}) {
	fmt.Printf("%T\n", v)
	str, ok := v.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Printf("传递的是一个字符串：%s\n", str)
	}
}

func justifyType(v interface{}) {
	fmt.Printf("%T\n", v)
	switch t := v.(type) {
	case string:
		fmt.Println("是一个 字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case bool:
		fmt.Println("是一个bool:", t)
	}
}

func main() {
	assign("100")
	justifyType("true")
}

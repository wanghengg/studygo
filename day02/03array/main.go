package main

import "fmt"

// 数组
// 存放元素的容器
// 必须指定存放的元素的类型和数组容量
// 数组的长度是数组类型的一部分

func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组的初始化
	// 如果不初始化，默认元素都是零值（布尔值：false，整型和浮点型都是0，字符串：""）
	fmt.Println(a1)
	fmt.Println(a2)

	// 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 初始化方式2
	// 根据初始化数组自动推断数组长度
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)
	fmt.Println(len(a10))

	// 初始化方式3
	// 根据索引初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历1
	cities := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// 数组的遍历2
	for i, v := range cities {
		fmt.Println(i, v)
	}
}

package main

import "fmt"

// 切片slice

func main() {
	// 切片的定义
	var s1 []int    // 定义一个存放int类型的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左包含右不包含（左闭右开）
	fmt.Println(s3)
	fmt.Printf("type of s3: %T\n", s3)
	fmt.Printf("type of a1: %T\n", a1)

	s6 := a1[3:] // {7, 9, 11, 13}
	fmt.Println(s6)
	a1[6] = 1300
	fmt.Println(s6) // 切片是引用类型，更改底层数组，切片也会更改
	fmt.Printf("len(s6):%d, cap(s6):%d\n", len(s6), cap(s6))

	// make函数构造切片
	b1 := make([]int, 5, 10) // 默认初始化零值
	fmt.Println(b1)
	fmt.Printf("b1=%v len(b1)=%d cap(b1)=%d\n", b1, len(b1), cap(b1))

	// 切片的赋值
	b3 := []int{1, 3, 5}
	b4 := b3
	fmt.Println(b3, b4)
	b3[0] = 100
	fmt.Println(b3, b4)

	for i, v := range b3 {
		fmt.Println(i, v)
	}
}

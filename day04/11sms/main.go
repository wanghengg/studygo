package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统可以查看/新增/删除学生
*/

type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student // 变量声明
)

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	// 遍历所有的学生
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	// 向allStudent中添加一个学生
	// 1. 创建一个学生
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号：")
	fmt.Scan(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scan(&name)
	// 2. 构造学生
	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

func deleteStudent() {
	fmt.Print("请输入要删除学生的学号：")
	var id int64
	fmt.Scan(&id)
	delete(allStudent, id)
}

func main() {
	allStudent = make(map[int64]*student, 48)
	// 1. 打印菜单
	fmt.Println("欢迎使用学生管理系统")
	fmt.Println(`
		1. 查看学生
		2. 新增学生
		3. 删除学生
	`)
	for {
		fmt.Print("请输入你要干什么：")
		// 2. 等待用户选择命令
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项!\n", choice)
		// 3. 执行对应函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("error")
		}
	}
}

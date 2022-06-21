package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周林",
		Age:  90000,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	// 反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}

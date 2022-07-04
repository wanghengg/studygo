package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// 明天的这个时间
	// 按照指定格式取解析一个字符串格式的时间，默认UTC+0
	parseTime, err := time.Parse("2006-01-02 15:04:05", "2019-08-04 14:41:50")
	if err != nil {
		fmt.Printf("parse time err: %v\n", err)
		return
	}
	fmt.Println(parseTime)

	// 按照UTC+8解析
	loc, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Printf("load loc failed, err: %v\n", err)
		return
	}
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-07-05 23:51:50", loc)
	if err != nil {
		fmt.Printf("parse time failed, err: %v\n", err)
		return
	}
	fmt.Println(timeObj)

	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

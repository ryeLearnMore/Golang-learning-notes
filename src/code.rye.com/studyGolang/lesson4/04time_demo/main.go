package main

import (
	"fmt"
	"time"
)

// 时间戳转化为时间格式
func timestamp2Timeobj(timestamp int64)  {
	timeObj := time.Unix(timestamp, 0) // 将时间戳转化为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 定时器
func tickDemo()  {
	ticker := time.Tick(time.Second) // 定义一个1s间隔的定时器
	// 设置每秒执行的任务
	for i := range ticker {
		fmt.Println(i) // 每秒都会执行的任务
	}
}

// 格式化打印时间
func formateDemo()  {
	/*
	2019-12-24 20:49
	2019/12/24 20:49
	20:49 2019/12/24
	2019/12/24
	*/
	now := time.Now()
	// 格式化的模板为Go的出生时间：2006年1月2号15点04分
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// 内置的time包
func main()  {
	// time.Time是一个struct
	now := time.Now()
	fmt.Printf("%#v\n", now) // time.Time{wall:0xbf789e5689585cec, ext:3997401, loc:(*time.Location)(0x54ca20)}
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// 时间戳
	fmt.Println(now.Unix()) // 1577190058
	fmt.Println(now.UnixNano())

	tm := now.Unix() + 100000
	timestamp2Timeobj(tm)

	// 定时器
	// tickDemo()
	
	//时间按格式输出
	formateDemo()
}

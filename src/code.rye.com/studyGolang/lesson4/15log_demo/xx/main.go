package main

// import (
// 	"fmt"
// )
import "code.rye.com/studyGolang/lesson4/15log_demo/mylog"

// 写了一个项目想要在代码中记录i日志
// 要使用mylog这个包
func main()  {
	f1 := mylog.NewFileLogger(mylog.DEBUG, "./", "test.log")
	// f1.Debug("这是一条测试的日志")
	// fmt.Println("可以申请IPO")
	userId := 10
	// f1.Info("这是一条debug级别的日志")
	f1.Debug("id是%d的用户一直在尝试登陆", userId)
}
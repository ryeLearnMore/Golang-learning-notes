package main

import "code.rye.com/studyGolang/lesson5/mylogger"

// 一个使用自定义日志库的用户程序
func main()  {
	logger := mylogger.NewFileLogger("./", "xxx.log")
	didi := "迪迪"
	logger.Debug("%s迪迪也爱我", didi)
}
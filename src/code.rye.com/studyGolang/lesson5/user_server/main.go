package main

import "code.rye.com/studyGolang/lesson5/mylogger"

// 一个使用自定义日志库的用户程序
func main()  {
	logger := mylogger.NewFileLogger("Info", "./", "xxx.log")
	didi := "迪迪"
	logger.Debug("%s迪迪也爱我", didi)
	logger.Info("Info 这是一条测试的日志")
	logger.Error("Error 这是一条测试的日志")
}
package main

// 测试mylogger的程序
import (
	"code.rye.com/studyGolang/lesson7/homework/mylogger"
)

var logger mylogger.Logger

func main()  {
	logger = mylogger.NewFileLogger("debug", "xx.log", "./")
	defer logger.Close()
	for {
		logger.Debug("下雨了")
		logger.Error("打雷了")
	}
}
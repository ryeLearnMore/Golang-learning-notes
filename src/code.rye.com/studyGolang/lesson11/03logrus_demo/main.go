package main

// logrus 示例

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 设置json格式
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"name": "didi",
		"age":  18,
	}).Warn("这是一条warnning级别的日志")
	logrus.Info("这是一条普通的日志")
}

package main

import (
	"fmt"
)

// 解析日志库的配置文件

// Config 是一个日志配置项

type Config struct {
	FilePath string `conf:"file_path" db:"name"`
	FileName string `conf:"file_name"`
	MaxSize  int64  `conf:"max_size"`
}

// 从conf文件中读取内容赋值给结构体指针
func parseConf(fileName string, interface{})  {
	
}

func main()  {
	// 1. 打开文件
	// 2. 读取内容
	// 3. 一行行读取内容，根据tag找结构体里面对应的字段
	// 4. 找到了要赋值
	parseConf()

}
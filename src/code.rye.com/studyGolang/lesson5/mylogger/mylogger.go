package mylogger

// 我的日志库文件

// level是一个自定义的类型，代表日志级别
type Level uint16

// 定义具体的日志级别常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)
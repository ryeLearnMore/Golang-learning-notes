package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 这是往文件里面写日志的代码

// FileLogger 文件日志结构体
type FileLogger struct {
	fileName string
	filePath string
	file *os.File
	errFile *os.File
}

// NewFileLogger 文件日志结构体构造函数
func NewFileLogger(fileName, filePath string) *FileLogger {
	fl := &FileLogger{
		fileName: fileName,
		filePath: filePath,
	}
	fl.initFile() // 根据上面的文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应的字段
	return fl
}

// 将指定的日志文件打开，赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	// 打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败，%v", logName, err))
	}
	f.file = fileObj

	// 打开错误日志的文件
	errLogName := fmt.Sprintf("%s.err", logName)
	// 打开文件
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败，%v", logName, err))
	}
	f.errFile = errFileObj
}

// 方法 
func (f *FileLogger) Debug(format string, args ...interface{})  {
	// f.file.Write()
	/*
	fmt.Fprintf() file
	fmt.Sprintf() string
	fmt.Errorf() error
	*/
	msg := fmt.Sprintf(format, args...)
	// 日志格式：[时间][文件：行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(2)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", 
		nowStr, fileName, line, funcName, "debug", msg)

	fmt.Fprintln(f.file, logMsg) // 利用fmt包将msg字符串写入f.file文件中
}
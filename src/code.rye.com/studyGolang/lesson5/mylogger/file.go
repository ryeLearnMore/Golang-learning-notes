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
	level Level
	fileName string
	filePath string
	file *os.File
	errFile *os.File
	maxSize int64 // 按照日志文件大小进行切分
}

// NewFileLogger 文件日志结构体构造函数
func NewFileLogger(levelStr, fileName, filePath string) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &FileLogger{
		level: logLevel,
		fileName: fileName,
		filePath: filePath,
		maxSize: 10 * 1024 * 1024,
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

// 检查是否需要拆分
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize // 当前进来的日志文件大小超过maxSize，就返回true
}

// 封装一个切分日志文件的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	// 判断当前日志文件的大小是否超过了maxSize
	// 切分文件
	fileName := file.Name() // 拿到文件的完整路径
	backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
	// 1. 把原来的文件关闭
	file.Close()
	// 2. 备份原来的文件
	os.Rename(fileName, backupName)
	// 3. 新建一个文件
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	return fileObj
}

// 将公用的记录日志的功能封装成一个单独的方法
func (f *FileLogger) log(level Level, format string, args ...interface{})  {
	if f.level > level {
		return 
	}
	// f.file.Write()
	/*
	fmt.Fprintf() file
	fmt.Sprintf() string
	fmt.Errorf() error
	*/
	msg := fmt.Sprintf(format, args...) // 得到用户记录的日志
	// 日志格式：[时间][文件：行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", 
		nowStr, fileName, line, funcName, logLevelStr, msg)
	// 往文件里面写之前要做一个检查
	if f.checkSplit(f.file) {
		f.file = f.splitLogFile(f.file)
	}
	fmt.Fprintln(f.file, logMsg) // 利用fmt包将msg字符串写入f.file文件中

	if level >= ErrorLevel {
		if f.checkSplit(f.errFile) {
			f.errFile = f.splitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, logMsg)
	}
}

// Debug 方法 
func (f *FileLogger) Debug(format string, args ...interface{})  {
	f.log(DebugLevel, format, args...)
	// if f.level > DebugLevel {
	// 	return 
	// }
	// // f.file.Write()
	// /*
	// fmt.Fprintf() file
	// fmt.Sprintf() string
	// fmt.Errorf() error
	// */
	// msg := fmt.Sprintf(format, args...) // 得到用户记录的日志
	// // 日志格式：[时间][文件：行号][函数名][日志级别] 日志信息
	// nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	// fileName, line, funcName := getCallerInfo(2)
	// logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", 
	// 	nowStr, fileName, line, funcName, "debug", msg)

	// fmt.Fprintln(f.file, logMsg) // 利用fmt包将msg字符串写入f.file文件中
}

// Info info方法Info
func (f *FileLogger) Info(format string, args ...interface{})  {
	f.log(InfoLevel, format, args...)
// 	if f.level > DebugLevel {
// 		return 
// 	}
// 	// f.file.Write()
// 	/*
// 	fmt.Fprintf() file
// 	fmt.Sprintf() string
// 	fmt.Errorf() error
// 	*/
// 	msg := fmt.Sprintf(format, args...)
// 	// 日志格式：[时间][文件：行号][函数名][日志级别] 日志信息
// 	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
// 	fileName, line, funcName := getCallerInfo(2)
// 	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", 
// 		nowStr, fileName, line, funcName, "info", msg)

// 	fmt.Fprintln(f.file, logMsg) // 利用fmt包将msg字符串写入f.file文件中
}

// Warn warn方法
func (f *FileLogger) Warn(format string, args ...interface{})  {
	f.log(WarningLevel, format, args...)
}

// Error error方法
func (f *FileLogger) Error(format string, args ...interface{})  {
	f.log(ErrorLevel, format, args...)
}

// Fatal fatal方法
func (f *FileLogger) Fatal(format string, args ...interface{})  {
	f.log(FatalLevel, format, args...)
}

// Close 关闭日志文件句柄
func (f *FileLogger) Close()  {
	f.file.Close()
	f.errFile.Close()
}
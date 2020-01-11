package main

import "fmt"
// 表示导入但不使用
import _ "code.rye.com/studyGolang/lesson4/02package/math_pkg"

// init()示例
var today = "星期天"

const Week = 7

type student struct {
	Name string
}

// 包被导入的时候会自动执行
func init()  {
	fmt.Println("包被初始化了")
	fmt.Println(Week) // 先打印”7“所以是先全局声明，再init，最后main()
}

func main()  {
	fmt.Println("这是main函数！")
}
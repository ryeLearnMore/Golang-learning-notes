package main

import "fmt"

// defer：延迟执行。作用：通常用于资源清理、文件关闭等
// 输出：
/*
start...
end...
3
2
1
*/
func main()  {
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
}
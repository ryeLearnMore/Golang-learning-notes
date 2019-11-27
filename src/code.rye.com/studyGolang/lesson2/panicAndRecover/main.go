package main

import "fmt"
// panic and recover, 常用于异常处理
/*
注：
1. recover()必须搭配defer使用
2. defer一定要在可能引发panic的语句之前定义
*/

func a()  {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error!")
		}
	}()
	panic("panic in b") // 不写defer语句会报错
}

func c() {
	fmt.Println("func c")
}

func main()  {
	/*
	func a
	func b error!
	func c
	*/
	a()
	b()
	c()
}
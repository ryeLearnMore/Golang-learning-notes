package main

import "fmt"
/*
声明方式共有四种，具体如下：
*/

func foo() (string, int) {
	return "alex", 900
}

func main() {
	// 1. 正常声明
	var name string
	var age int
	// 声明变量必须要使用
	fmt.Println(name)
	fmt.Println(age)

	// 2. 批量声明
	var (
		a string // ""
		b int // 0
		c bool // false
	)
	fmt.Println(a, b, c)

	var x string = "学习"
	fmt.Println(x)
	// Printf格式化打印：欢迎%s登陆
	fmt.Printf("%s嘿嘿嘿\n", x)

	//3. 类型推导（编译器根据变量初始值的类型，指定给定变量）
	var y = 200
	var z = true
	fmt.Println(y, z)

	// 4. 短变量声明（只能在函数内部使用）
	nazha := "嘿嘿嘿" // 相当于var nazha string = "嘿嘿嘿"
	fmt.Println(nazha)

	// 调用foo函数， _ （匿名变量)用于接收不需要的变量值
	aa, _ := foo()
	// fmt.Println(aa, bb)
	fmt.Println(aa)
}
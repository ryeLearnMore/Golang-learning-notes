package main

import "fmt"

// 函数进阶之变量作用域

// 定义全局变量num
var num = 10

// 定义函数
func testGlobal()  {
	// 可以再函数中访问全局变量
	fmt.Println("全局变量：", num)
}

func testPrivate()  {
	num := 100
	// name := "haha"

	// 可以在函数中使用变量
	// 1. 先在自己函数中查找，找到了就用自己函数的
	// 2. 函数中找不到变量就从外层找全局变量
	fmt.Println("变量是：", num)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main()  {
	// point 1
	// testGloba l() // 全局变量： 10

	// point 2
	// 外层不能访问函数内层的变量（局部变量）
	// fmt.Println(name)

	// point 3
	testPrivate() // 变量是： 100

	// point 4
	// 函数可以作为变量
	abc := testGlobal
	fmt.Printf("%T\n", abc) // func()
	abc()

	// point 5
	// 函数可以作为参数
	r1 := calc(100, 200, add)
	fmt.Println(r1) // 300
	r2 := calc(100, 200, sub)
	fmt.Println(r2) // 100
}
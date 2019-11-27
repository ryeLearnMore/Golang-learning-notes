package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包

// 匿名函数
func anonymousFunc() {
	// point1: 匿名函数执行方式1
	sayHello := func ()  {
		fmt.Println("匿名函数执行方式1")
	}
	sayHello()

	// point2: 匿名函数执行方式2
	func ()  {
		fmt.Println("匿名函数执行方式2")
	}()
}

// 闭包
// 定义一个函数它的返回值是一个函数，即把函数作为返回值
func a() func() {
	return func ()  {
		fmt.Println("闭包输出")
	}
}

// 闭包优化1
func a1() func() {
	name := "带有变量的闭包"
	return func ()  {
		fmt.Println("闭包输出: ", name)
	}
}

// 闭包优化2
func a2(name string) func() {
	return func ()  {
		fmt.Println("闭包输出: ", name)
	}
}

// 闭包的应用1：使用闭包做后缀名检测
// suffix: 后缀；词尾
func makeSuffixFunc(suffix string) func(string) string {
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包应用2：
func calc(base int) (func(int) int, func(int) int) {
	add := func (i int) int {
		base += i
		return base
	}
	
	sub := func (i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 匿名函数
	// anonymousFunc()

	r := a()
	r() // 相当于执行了a函数内部的匿名函数

	// 闭包 = 函数 + 外层变量的引用
	r1 := a1()
	r1()

	r2 := a2("形参")
	r2()

	r3 := makeSuffixFunc(".txt")
	ret := r3("娜扎")
	fmt.Println(ret)

	r4 := makeSuffixFunc(".avi")
	ret1 := r4("haha")
	fmt.Println(ret1)

	// 判断一个函数是不是闭包，就要看内部函数是否引用外部函数的变量
	x, y := calc(100)
	ret11 := x(200)
	fmt.Println(ret11) // base = 100 + 200

	ret22 := y(200)
	fmt.Println(ret22) // base = 300 - 200
}
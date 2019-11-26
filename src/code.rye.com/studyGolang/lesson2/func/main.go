package main

import "fmt"

// 定义一个不需要参数也没有返回值的函数，sayHello
func sayHello()  {
	fmt.Println("Hello!")
}

func sayHello2(name string)  {
	fmt.Println("Hello ", name)
}

func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intSum2(a int, b int) (ret int) {
	ret = a + b
	return ret // 或者可以不写ret，直接return
}

// 函数接收可变参数,在参数后面加...表示可变参数
// 可变参数在函数中是切片类型
func intSum3(a ...int) int {
	// fmt.Println(a) // []
	// fmt.Printf("%T\n", a) // []int
	ret := 0
	for _, arg := range a{
		ret += arg
	}
	return ret
}

// 注：固定参数和可变参数同时出现时，可变参数要放在最后
// go语言的函数中没有默认参数
func intSum4(a int, b ...int) int {
	// fmt.Println(a) // []
	// fmt.Printf("%T\n", a) // []int
	ret := a
	for _, arg := range b{
		ret += arg
	}
	return ret
}

// Go语言中函数参数类型简写
func intSum5(a, b int) int {
	ret := a + b
	return ret
}

// 定义具有多个返回值的函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main()  {
	// 函数调用
	// sayHello()

	// name := "rye"
	// sayHello2(name)

	// r := intSum(10, 20)
	// fmt.Println(r)

	// fmt.Println(intSum3(10, 20, 3, 4))
	// fmt.Println(intSum4(1, 20, 3, 4)) 
	x, y := calc(100, 200)
	fmt.Println(x, y)
}
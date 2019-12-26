package main

import "fmt"

// type nullInterface interface {} 太繁琐

func ShowType(x interface{})  {
	// 因为我这个函数可以接收任意类型的变量
	// 类型断言
	v1, ok := x.(int)
	if !ok {
		// 说明猜错了
		fmt.Println("x不是int类型")
	} else {
		fmt.Println("x就是一个int类型", v1)
	}

	v2, ok := x.(string)
	if !ok {
		fmt.Println("不是string")
	} else{
		fmt.Println("x是一个string类型", v2)
	}
}

func justifyType(x interface{})  {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int %v\n", v)
	case bool:
		fmt.Printf("x is a bool %v\n", v)
	default:
		fmt.Printf("unsupport type!")
	}
}

func main()  {
	ShowType(100)
	ShowType("haha")

	// switch版本
	justifyType(100)
}
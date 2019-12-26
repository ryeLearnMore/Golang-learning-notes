package main

import "fmt"

// 空接口
// 作用：使用空接口可以接收任意类型的函数参数

func showType(a interface{})  {
	fmt.Printf("type:%T\n", a)
}


func main()  {
	var x interface {}
	x = 100
	fmt.Println(x)
	x = "zhangdi"
	fmt.Println(x)
	x = false
	fmt.Println(x)
	x = struct{name string}{name: "张迪"}
	fmt.Println(x)

	showType(x) // type:struct { name string }

	// 定义一个值为空接口的map，那么就可以存储任意类型的值
	var stuInfo = make(map[string]interface{}, 100)
	stuInfo["zhangdi"] = 100
	stuInfo["asdfa"] = true
	stuInfo["asdf"] = 88.88
	stuInfo["aasg"] = "zhangdi"
	fmt.Println(stuInfo)
}
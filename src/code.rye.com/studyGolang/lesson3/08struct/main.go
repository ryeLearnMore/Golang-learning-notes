package main

import "fmt"

type student struct {
	name string
	age int8
}

func main()  {
	var stu1 = student{
		name: "桥本环奈",
		age: 18,
	}

	stu2 := stu1 //将结构体stu1的值完整的复制一份给了stu2
	stu2.name = "长泽雅美"
	fmt.Println(stu1.name) // 桥本环奈
	fmt.Println(stu2.name) // 长泽雅美

	stu3 := &stu1 // 将stu1对应的地址赋值给了stu3，stu3的类型是一个*student
	fmt.Printf("%T\n", stu3) // *main.student
	stu3.name = "深田恭子"
	fmt.Println(stu1.name, stu2.name, stu3.name) // 深田恭子 长泽雅美 深田恭子
}
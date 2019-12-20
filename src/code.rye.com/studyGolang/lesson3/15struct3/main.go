package main

import "fmt"

// 结构体的嵌套
type address struct {
	province string
	city string
}

type student struct {
	name string
	age int
	// addr address // 此处嵌套了别的结构体
	address
}

func main()  {
	var stu1 = student{
		name:"zhangdi",
		age:18,
		address: address{
			province: "Liaoning",
			city: "beipiao",
		},
	}
	fmt.Println(stu1)
	fmt.Println(stu1.name)
	// fmt.Println(stu1.addr.province)
	fmt.Println(stu1.province) // 匿名字段支持简写可直接访问，但当匿名字段有冲突时，必须显示写
}


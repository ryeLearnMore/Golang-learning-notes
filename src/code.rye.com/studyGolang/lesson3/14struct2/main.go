package main

import "fmt"

// 匿名字段
type student struct {
	name string
	string
	int
}

func main()  {
	var stu1 = student{
		name: "haojie",
		string: "asdf",
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.string)
	fmt.Println(stu1.int)
	/*
	haojie
	asdf
	0
	*/
}
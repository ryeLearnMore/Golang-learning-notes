package main

import (
	"fmt"
	"reflect"
)

// 反射的TypeOf()
func reflectType(x interface{})  {
	t := reflect.TypeOf(x) // 拿到x的动态类型信息
	fmt.Printf("type: %v\n", t)
	// fmt.Printf("%T\n", x) // 能打印出原本类型的原理就是用的反射

	fmt.Printf("name: %v, kind: %v\n", t.Name(), t.Kind())
	// 指针类型的t.Name()是空
}

type cat struct {
	name string
}

type person struct {
	name string
	age int
}

func main()  {
	reflectType(100) // type: int
	reflectType("didi") // type: string
	reflectType(false) // type: bool
	
	c := cat{
		name: "didi",
	}
	reflectType(c) // type: main.cat  name: cat, kind: struct

	p := person{
		name: "didi",
		age: 18,
	}
	reflectType(p) // type: main.person  name: person, kind: struct

	var i int32 = 100
	var f float32 = 12.34
	reflectType(&i) // type: *int32  name: , kind: ptr
	reflectType(&f) // type: *float32  name: , kind: ptr
}
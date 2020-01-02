package main

import (
	"fmt"
	"reflect"
)

func modifyValue(x interface{})  {
	v := reflect.ValueOf(x) // v的类型是reflect.Value
	kind := v.Kind()
	fmt.Println(kind)
	if kind == reflect.Ptr {
		// 传入的是一个指针
		// 取到指针对应的值再去修改
		v.Elem().SetInt(1234)
	}
}

func main()  {
	var a int64 = 100
	// 将a传入一个函数，在函数中修改a的值，该函数可以接收任意类型的值
	modifyValue(&a) // ptr
	fmt.Println(a) // 1234
}
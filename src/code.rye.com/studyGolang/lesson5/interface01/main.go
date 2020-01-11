package main

// 接口值，由两部分组成：类型和值
// 下面介绍了断言
// 类型太多了，类型断言猜不全，所以通过反射直接拿到接口的值的动态类型和动态值

import "fmt"

func main()  {
	var x interface{} // <Type, Value>
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1
	x = a // <int64, 100>
	x = b // <int32, 10>
	x = c // <int8, 1>
	x = 12.34 // <float64, 12.34>
	x = false // <bool, false>
	fmt.Println(x)

	// 类型断言（猜这个值的类型）
	// 如果猜对了，ok = true, value = 对应的值
	// 如果猜错了，ok = false, value = 对应类型的零值
	value, ok := x.(string)
	fmt.Printf("ok:%t, value:%#v, value type: %T\n", ok, value, value)
}
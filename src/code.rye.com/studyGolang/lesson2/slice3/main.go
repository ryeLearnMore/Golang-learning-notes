package main

import "fmt"
//切片是引用类型

func main()  {
	a := []int{1, 2, 3}
	b := a // 直接赋值
	b[0] = 100
	fmt.Println(a) // [100 2 3]
	fmt.Println(b) // [100 2 3]， 结果相等，说明为引用型，与数组中的值类型不同

	// 深拷贝
	var c []int // 只是声明，还没有申请内存
	c = make([]int, 3, 3)
	copy(c, a)

	a[0] = 1
	fmt.Println(a) // [1 2 3]
	fmt.Println(c) // [100 2 3]
}
package main

import "fmt"

func main()  {
	// 定义一个数组
	a := [...]int{1, 3, 5, 7, 9, 11, 13}
	// 基于数组得到一个切片
	b := a[:]
	// 修改切片中的第一个元素为100
	b[0] = 100
	// 打印数组中第一个元素的值是多少？
	fmt.Println(a[0]) // 是100，不是1，因为切片是基于数组的

	// 问切片的长度是多少？容量是多少？
	c := a[2: 5]
	fmt.Println(c) // [5 7 9]
	fmt.Println(len(c)) // 3
	fmt.Println(cap(c)) // 5

	// 问切片的长度是多少？容量是多少？
	d := c[:5]
	fmt.Println(d) // [5 7 9 11 13]
	fmt.Println(len(d)) // 5
	fmt.Println(cap(d)) // 5

	// 问切片的长度是多少？容量是多少？
	e := d[2:]
	fmt.Println(e) // [9 11 13]
	fmt.Println(len(e)) // 3
	fmt.Println(cap(e)) // 3

}
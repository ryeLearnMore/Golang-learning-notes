package main

import "fmt"

func main()  {
	var a = [3]int{1, 2, 3} // 数组，有固定长度
	// 声明切片的方式1：直接声明
	var b = []int{1, 2, 3} // 切片
	fmt.Println(a, b)
	fmt.Printf("a: %T, b: %T\n", a, b)

	/*
	// 从切片中取值
	fmt.Println(b[1])

	// 声明切片的方式2： 从数组中得到切片
	var c []int
	c = a[0:3] // c = a[:]
	fmt.Printf("c: %T\n", c)
	fmt.Println(c)

	// 冒号切 左包含右不包含
	d := a[:2] // 从开始切到索引为2（不包含2）
	e := a[1:] // 从索引为1 开始切到最后
	fmt.Println(d)
	fmt.Println(e)
	*/

	// 切片的大小（目前元素的数量）
	fmt.Println(len(b))
	// 容量（底层数组最大能放多少元素）
	x := [...]string{"北京", "shanghai", "shenzhen", "guangzhou"}
	y := x[1: 3]
	fmt.Println(y)
	fmt.Println("切片y的长度是：", len(y)) // 2
	fmt.Println("切片y的容量是：", cap(y)) // 3
	
}
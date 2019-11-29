package main

import "fmt"
/*
new 和 make的区别
new: 是用来初始化值类型指针的
make：是用来初始化slice、map、channel
*/
func main()  {
	/*
	错误写法：
	var b *string
	var c *[3]int
	var a *int // a是一个int类型的指针
	*a = 10 // 会报错。因为上面只是声明地址，但并未创建内存空间
	*/
	var a = new(int) // 得到一个int类型的指针
	fmt.Println(a)
	*a = 10
	fmt.Println(*a)

	var c = new([3]int)
	fmt.Println(c)
	c[0] = 1
	fmt.Println(*c)
}
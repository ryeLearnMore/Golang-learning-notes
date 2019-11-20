package main

import "fmt"
import "math"

func main()  {
	// 10进制
	var a int = 10 
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010, 占位符%b表示二进制

	// 8进制，以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77
	
	// 16进制， 以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 变量的内存地址
	fmt.Printf("%p \n", &a) // 0xc00004c080，占位符%p表示十六进制的内存地址

	// 以10进制打印8进制数
	fmt.Println(a, b) //10, 63

	floatNumber()
}

func floatNumber() {
	fmt.Println(math.MaxFloat64)
	// 字符串多行输出
	var s1 = 
`
这
	是一个
	多行文本
	`
	fmt.Println(s1)
}
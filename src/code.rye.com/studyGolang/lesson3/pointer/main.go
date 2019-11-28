package main

import "fmt"
/*
指针和地址有什么区别？
地址：就是内存地址（用字节来描述的内存地址）
指针：指针是带属性的（所以字符串类型的指针不能赋值给int型的指针）

&和*
&: 表示取地址
*：根据地址取值
Go的指针和C++的不太相同。Go的指针只能取地址，不能更改地址
*/

func main() {
	var a int
	fmt.Println(a)
	b := &a // 取变量a的内存地址
	fmt.Println(b)
	fmt.Printf("b = %v\n", b)
	c := "haojie"
	fmt.Printf("&c=%v\n", &c)
	// b = &c // 取c的内存地址不能赋值给b，因为两者的类型不同（一个是string, 一个是int），若相同就可以了

	d := 100
	b = &d
	fmt.Println(b)
	// *取地址对应的值
	fmt.Println(*b)


	// 指针的应用
	a11 := [3]int{1, 2, 3}
	modifyArray(a11) // 在函数中复制了数组赋值给内部的a1，并没有对原数组进行更改
	fmt.Println(a11) // [1 2 3]，不是100，2，3

	modifyArray2(&a11)
	fmt.Println(a11)
}

// 定义一个修改数组第一个元素为100的函数
func modifyArray(a1 [3]int)  {
	a1[0] = 100 // 只是修改内部a1这个数组
}

// 定义一个修改数组第一个元素为100的函数
// 接收的参数是一个数组的指针
func modifyArray2(a1 *[3]int)  {
	// 写法1
	// (*a1)[0] = 100 // 只是修改内部a1这个数组
	// 写法2：语法糖：因为Go语言中指针不支持修改
	a1[0] = 100 
}
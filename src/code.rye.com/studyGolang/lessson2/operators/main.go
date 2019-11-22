package main

import "fmt"

func main()  {
	//算术运算符
	n1 := 13
	n2 := 3
	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2) // 4
	fmt.Println(n1 % n2) // 1

	// 关系运算符
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)

	// 逻辑运算符
	a := true
	b := false
	//两个条件都成立
	fmt.Println(a && b)
	// 两个条件有一个成立为真
	fmt.Println(a || b)
	//原来为真，取非为假。反之亦然
	fmt.Println(!a)

	// 位运算符
	// & | ^ << >>


}
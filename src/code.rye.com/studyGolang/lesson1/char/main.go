package main

import "fmt"

func main()  {
	s1 := "Golang"
	c1 := 'G' // ascii码下占一个字节（8位）
	fmt.Println(s1, c1) // Golang 71
	s2 := "中国"
	c2 := '中' // utf-8编码中一个中文占三个字节
	fmt.Println(s2, c2) // 中国 20013

	s3 := "hello沙河"
	fmt.Println(len(s3)) // 5+3+3

	/*
	1. uint8类型，也叫byte类型，代表ASCII码的一个字符
	2. rune类型，代表一个UTF-8字符，当处理中文或日文等符合字符需要用到rune类型，rune类型实际上是一个int32
	*/

	//遍历字符串
	for i := 0; i < len(s3); i ++ {
		fmt.Printf("%c ", s3[i])
	}
	fmt.Println()
	// for range循环，是按照rune类型去遍历的
	for _, v := range s3 {
		fmt.Printf("%c ", v)
	}

	fmt.Println()
	// 强制类型转换
	s5 := "big"
	// 将字符串强制转换成字节数组类型
	byteArray := []byte(s5)
	fmt.Println(byteArray) // [98 105 103]
	byteArray[0] = 'p'
	//将字节数组强制转换成字符串类型
	s5 = string(byteArray)
	fmt.Println(s5)
}
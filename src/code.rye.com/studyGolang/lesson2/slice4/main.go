package main

import "fmt"

// 切片中删除某个元素的方法
func main()  {
	a := []string{"beijing", "shanghai", "guangzhou"}
	a = append(a, "sichuan")
	a = append(a[:1], a[2:]...)
	fmt.Println(a)
}
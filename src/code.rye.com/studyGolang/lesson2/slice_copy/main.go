package main

import "fmt"

func main()  {
	// 切片的copy
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5)
	c := b
	copy(b, a) // 将切片a复制到切片b(深拷贝)
	b[0] = 100
	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(b) // [100 2 3 4 5]
	fmt.Println(c) // [100 2 3 4 5]
}
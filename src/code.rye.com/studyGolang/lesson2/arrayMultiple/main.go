package main

import "fmt"

// 多维数组
func main() {
	var a [3]int
	a = [3]int{1, 2, 3}

	// 声明二维数组
	var b [3][2]int
	// 赋值
	b = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println(a)
	fmt.Println(b) 

	// 声明的同时完成赋值
	var c = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6}, // , 不可以省略
	}
	fmt.Println(c)

	// 注意事项：多维数组除了第一层其他层都不能用 ...
	var d = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6}, // , 不可以省略
	}
	fmt.Println(d)

	// 多维数组的使用
	fmt.Println(d[1][0])
	// 多维数组的遍历，方法1
	for i := 0; i < len(d); i ++ {
		for j := 0; j < len(d[i]); j ++ {
			fmt.Printf("%d-%d\n", i, d[i][j])
		}
	}
	// 多维数组的遍历，方法2
	for _, v1 := range d {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型，不是引用类型
	a1 := [2]int{1, 2}
	b1 := a1 // 把a整个拷贝一份赋值给了b
	b1[0] = 100
	fmt.Println(a1) // [1, 2]
	fmt.Println(b1) // [100, 2]

	// 二维数组
	c1 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	d1 := c1 // 复制的时候都是直接拷贝，属于值类型
	d1[0][0] = 100
	fmt.Println(c1) // [[1 2] [3 4] [5 6]]
	fmt.Println(d1) // [[100 2] [3 4] [5 6]]
}
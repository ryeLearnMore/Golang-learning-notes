package main

import "fmt"
// 1. 求数组中所有元素的和
func getSum()  {
	a1 := [...]int{1, 3, 5, 7, 8}

	// 数组求和
	sum := 0
	// :=相当于声明变量且赋值
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)
}

// 2. 找出数组中和为指定值的两个元素下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素下标，分别为(0, 3), (1, 2)
func getIndex()  {
	// 思路：遍历数组，依次取出每个元素，计算一下other=8-当前值是否在数组中
	a1 := [...]int{1, 3, 5, 7, 8}
	for index, value := range a1 {
		other := 8 - value
		for index2 := index + 1; index2 < len(a1); index2 ++ {
			if a1[index2] == other {
				fmt.Printf("结果是：(%d, %d)\n", index, index2)
			}
		}
	}
}

func main()  {
	// getSum()
	getIndex()
}
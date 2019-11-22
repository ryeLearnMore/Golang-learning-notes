package main

import "fmt"
// 打印99乘法表
func chengFaBiao()  {
	for i := 1; i < 10; i ++ {
		for j := 1; j <= i; j ++ {
			fmt.Printf("%d*%d=%d  ", i, j, i * j)
		}
		fmt.Println()
	}
}

// 打印200-1000的素数
func suShu()  {
	for i := 200; i < 1001; i ++ {
		flag := true
		for j := 2; j < i/2; j ++ {
			if i % j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%d是素数\n", i)
		}
	}
}

func main() {
	// chengFaBiao()
	suShu()
}
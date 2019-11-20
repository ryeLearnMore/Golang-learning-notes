package main

import "fmt"
// goto

func oldGoTO() {
	flag := false
	for i := 0; i < 5; i ++ {
		for j := 0; j < 3; j ++ {
			if i == 2 && j == 2 {
				flag = true
				break
			}
			fmt.Printf("%d...%d\n", i, j)
		}

		if flag {
			break
		}
	}
}

func newGoTO() {
	// flag := false
	for i := 0; i < 5; i ++ {
		for j := 0; j < 3; j ++ {
			if i == 2 && j == 2 {
				// flag = true
				// break
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%d...%d\n", i, j)
		}

		// if flag {
		// 	break
		// }
	}
	breakTag:
	fmt.Println("两层for循环结束")
}
func main()  {
	oldGoTO()
	newGoTO()
}
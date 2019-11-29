package main

import "fmt"

func main()  {
	defer func() {
		err := recover() // 尝试将函数从当前异常状态回复
		fmt.Println("recover抓到了panic异常：", err)
	}()
	var a []int
	a[0] = 100 // 此处会引发panic

	fmt.Println("不会执行此句")
}
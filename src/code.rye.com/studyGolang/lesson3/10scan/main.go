package main

import "fmt"

// 从终端获取用户的输入内容
func main()  {
	var(
		name string
		age int
		married bool
	)
	fmt.Println(name, age, married)
	// 1. 输入方式1：fmt.Scan()
	// fmt.Scan(&name, &age, &married)
	// 2. 输入方式2：fmt.Scanf():按照固定格式输入才行
	// fmt.Scanf("name:%s age:%d married:%t\n", &name, &age, &married)
	// 3. 输入方式3：fmt.Scanln()：输入检测到了换行就立即结束
	fmt.Scanln(&name, &age, &married)
	fmt.Println(name, age, married)
}
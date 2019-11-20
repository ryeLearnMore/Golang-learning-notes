package main

import "fmt"

func main()  {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("澳门首家赌场开业了！")
	// }else if age < 18 {
	// 	fmt.Println("Warning...")
	// }else {
	// 	fmt.Println("成年了")
	// }

	// if的特殊用法
	age := 18
	if age == 18 {
		fmt.Println("成年了")
	}

	if age2 := 28; age2 > 18 {
		fmt.Println("成年了")
	}

	// fmt.Println(age2) 出错。因为第二个if判断语句中，age2的作用域只在if循环里
}
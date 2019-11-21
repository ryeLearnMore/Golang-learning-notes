package main

import "fmt"

func main()  {
	// 1. 先定义，然后赋值
	var a [3]int // 定义一个长度为5，存放int类型的数组
	var b [5]int
	// 初始化
	a = [3]int{1, 2, 3}
	b = [5]int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(b)

	// var c [3]string = [3]string{"北京", "上海", "广州"}
	// 上面的语句语法没问题，但是略冗余，因为等号后面的会推断出是数组类型
	var c = [3]string{"北京", "上海", "广州"}
	fmt.Println(c)

	// 2.  ...表示让编译器去数一下有多少个初始值，然后给变量赋值
	var d = [...]int{1, 23, 3, 345}
	fmt.Printf("c: %T, d: %T\n", c, d)

	// 3. 根据索引值初始化
	var e [5]int
	e = [5]int{3: 1}
	fmt.Println(e)


	// 数组的基本使用
	fmt.Println(e[3])
	// 遍历数组的方式1
	for i := 0; i < len(e); i ++ {
		fmt.Printf("%d ", e[i])
	}
	fmt.Println()
	// 遍历数组的方式2: for range循环
	for _, value := range e{
		fmt.Println(value)
	}

}
package main

import "fmt"
// 通用占位符
/*
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	打印值的类型
%%	百分号
*/
func main()  {
	var a = 100
	var b = "哈哈"
	var c = false

	fmt.Println(a, b, c)
	// %v俗称占位符
	fmt.Printf("a=%v\n", a)
	fmt.Printf("a的类型是%T\n", a)
	// %%转义
	fmt.Printf("%d%%", a)

	fmt.Printf("|%d|\n", a)
	fmt.Printf("|%8d|\n", a)
	fmt.Printf("|%-8d|\n", a)
	fmt.Printf("|%08d|\n", a)

	// 浮点数
	f1 := 3.141592654
	fmt.Printf("%f\n", f1) // 3.141593
	// 保留小数点后两位
	fmt.Printf("%.2f\n", f1) // 3.14
	// 保留小数点总共两位
	fmt.Printf("%.2g\n", f1) // 3.1

	// 字符串
	s1 := "这是一个字符串\""
	fmt.Printf("s1: %s\n", s1)
	fmt.Printf("s1: %q\n", s1)

}
package main

import "fmt"

func main() {
	a := [...]int{1, 2} // a为数组，属于值类型
	b := a[:] // b为切片，引用类型
	/*
	输出结果一样
	*/
	fmt.Printf("%p\n", &a) // 找变量a的地址
	fmt.Printf("%p\n", b) // 找变量b指向的内存地址

	// 扩容策略
	/*
	1. 每次只追加一个元素，每一次就是上一次的2倍
	2. 追加的超过原来的1倍，就等于原来的容量+扩容元素的个数最接近的偶数
	3. 如果切片的容量大于1024，后续就每次扩容0.25倍
	*/
	fmt.Println("b的容量：", cap(b))
	b = append(b, 3, 4, 5, 6, 7, 8)
	fmt.Println("b的容量：", cap(b))
	b = append(b, 8)
	fmt.Println("b的容量：", cap(b))

	fmt.Println(a, b) // [1 2] [1 2 3 4 5 6 7 8 8]，数组本身并未因切片而发生变化
	// 切片的总容量如果超过1024，再追加0.25倍

	makeUsage()
}

// make：用来给引用类型做初始化（申请内存空间）
// new：用来创建值类型

// 用法
/*
make用于给引用类型申请内存空间
切片必须初始化或用append才能使用
var s []int
s = make([]int, len, cap)
对切片初始化时，容量尽量大一点，减少程序运行期间再去动态扩容
*/
func makeUsage()  {
	var s1 []int
	// s1[0] = 100 // 不能这么做
	// s1 = append(s1, 100) // 可以这么用
	// fmt.Println(s1)
	s1 = make([]int, 3, 5)
	fmt.Println(s1)
}
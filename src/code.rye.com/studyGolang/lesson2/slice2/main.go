package main

import "fmt"

func main() {
	/*
	结果：
	a: [], len: 0, cap: 0, ptr: 0x55b988
	a: [1], len: 1, cap: 1, ptr: 0xc04204e090
	a: [1 1], len: 2, cap: 2, ptr: 0xc04204e0b0
	a: [1 1 1], len: 3, cap: 4, ptr: 0xc04204c0c0
	a: [1 1 1 1], len: 4, cap: 4, ptr: 0xc04204c0c0
	a: [1 1 1 1 1], len: 5, cap: 8, ptr: 0xc042072080
	*/

	// 切片三要素：
	/*
	1. 地址：切片中第一个元素指向内存空间!.需要理解，可参考slice5.go中的例子理解。简而言之，跟在一个数组上切几次没有关系，跟第一个切的位置有关
	2. 大小：切片中目前元素的个数 len()
	3. 容量：底层数组最大能存放的元素个数 cap()
	
	注：
	切片支持自动扩容，扩容策略是每一次都是上一次的2倍
	a = append(a, element)
	*/
	var a = []int{}
	fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a, len(a), cap(a), a) // 地址没变
	a = append(a, 1)
	fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a, len(a), cap(a), a) // 地址没变
	a = append(a, 1)
	fmt.Printf("a: %v, len: %d, cap: %d, ptr: %p\n", a, len(a), cap(a), a)
}
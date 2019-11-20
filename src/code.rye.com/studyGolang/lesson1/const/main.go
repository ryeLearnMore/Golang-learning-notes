package main
import "fmt"
/*
iota总结：
0. const声明如果不写，就默认和上一行一致
1. 遇到const iota就初始化为0
2. const中每新增一行变量声明iota就递增1
*/
const pi = 3.14

// 批量声明常量
const (
	a = 100
	b = 1000
	c // 1000，不是0 相当于c = 1000
	d // 同上
)

// iota枚举
const (
	aa = iota // 0
	bb        // 1
	cc		  // 2
	dd        // 3
)

// iota应用，每新增一行，就累加一次
//iota应用1
const (
	n1 = iota // 0
	n2        // 1
	_         // 省略
	n4        // 3
)
//iota应用2
const (
	nn1 = iota // 0
	nn2 = 100  // 100
	nn3 = iota // 2，此处不是1
	nn4        // 3，此处不是2
)
// iota应用3
const (
	_ = iota // 0
	KB = 1 << (10 * iota) // 1 << 10 --> 2的10次方 = 1024
	MB = 1 << (10 * iota) // 1 << 20
	GB = 1 << (10 * iota) // 1 << 30
	TB = 1 << (10 * iota) // 1 << 40
	PB = 1 << (10 * iota) // 1 << 50
)

//iota应用4
const (
	x1, x2 = iota + 1, iota + 2 // iota = 0; x1 = 1, x2 = 2
	x3, x4 // 相当于x3, x4 = iota + 1, iota + 2. 此时，iota = 1; x3 = 2, x4 = 3
	x5, x6// 相当于x5, x6 = iota + 1, iota + 2. 此时，iota = 1; x5 = 3, x6 = 4
)
//常量
func main()  {
	// pi = 3.1415926 // 常量不让赋值
	fmt.Println(pi)
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc, dd)
	fmt.Println(n1, n2, n4)
	fmt.Println(nn1, nn2, nn3, nn4)
	fmt.Println(x1, x2, x3, x4, x5, x6)
}
package main
// 04也是panic—recover，内容相似，就没记录
import "fmt"

// 自定义类型。NewInt是一个新类型
type NewInt int

// 类型别名：只存在代码编写过程中，代码编译之后根本不存在MyInt
// 作用：提高代码可读性
// 两者区别见输出
type MyInt = int

func main()  {
	var a NewInt
	fmt.Println(a)
	fmt.Printf("%T\n", a) // main.NewInt
	
	var b MyInt
	fmt.Println(b)
	fmt.Printf("%T\n", b) // int
}
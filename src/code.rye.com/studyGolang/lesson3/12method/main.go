package main
import "fmt"
// 函数和方法的区别
// 函数是谁都可以调用的
// 方法就是某个具体的类型才能调用的函数
type people struct {
	name string
	gender string
}

// 方法：函数指定接受者之后就是方法
// 在go语言中约定俗成不用this，也不用self，而是用后面类型的首字母的小写
func (p *people) dream()  {
	p.gender = "lady" // 不加指针，输出的是女，不是lady。这是因为结构体都是值类型。
	fmt.Printf("%s的梦想是不用上班也有钱\n", p.name)
}

func main()  {
	var zhangdi = people {
		name: "张迪",
		gender: "女",
	}
	// (&zhangdi).dream()
	// 语法糖
	zhangdi.dream()
	fmt.Println(zhangdi.gender)
}

/*
总结：
什么时候应该使用指针类型（如接受者是否用指针）
1. 需要修改接受者中的值
2. 接受者是拷贝代价比较大的对象
3. 一般情况下通常都采用指针接受者
*/
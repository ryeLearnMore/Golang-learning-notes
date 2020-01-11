package main

import "fmt"

//接口实现一个洗衣机
// 只要一个类型它实现了wash()和dry()方法，我们就称这个类型实现了xiyiji这个接口
type xiyiji interface {
	wash()
	dry()
}

type Haier struct{
	name string
	price float64
	mode string
}

type xiaotiane struct{
	name string
	price float64
	mode string
}

func (h Haier) wash()  {
	fmt.Println("海尔洗衣机能洗衣服")
}

func (h Haier) dry()  {
	fmt.Println("海尔洗衣机能自带甩干")
}

func main()  {
	// var haier xiyiji
	// fmt.Println(haier)
	var a xiyiji // 声明一个xiyiji类型的变量a
	h1 := Haier {
		name : "zhangdi",
		price: 999,
		mode: "滚筒",
	}
	fmt.Printf("%T\n", h1) // main.Haier
	a = h1 // 接口是一种类型，一种抽象的类型
	fmt.Println(a) // {zhangdi 999 滚筒}
	a.wash()


}

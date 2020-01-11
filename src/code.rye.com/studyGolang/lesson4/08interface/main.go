package main

import "fmt"

// 接口嵌套
// 定义一个会说话的接口
type speaker interface {
	speak()
}

// 定义一个会移动的接口
type mover interface {
	move()
}

// 此处嵌套接口
type animal interface {
	speaker
	mover
}

type cat struct {
	name string
}

func (c cat) speak()  {
	fmt.Println("喵喵喵")
}

func (c cat) move()  {
	fmt.Println("猫会动")
}

func main()  {
	var x animal
	x = cat{name: "zhangdi"}
	x.move()
	x.speak()
}
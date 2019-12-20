package main

import "fmt"

// 结构体内嵌模拟“继承”

type animal struct {
	name string
}

// 定义一个动物会动的方法
func (a *animal) move()  {
	fmt.Printf("%s会动\n", a.name)
}

// 定义一个狗的结构体
type dog struct {
	feet int
	*animal // 嵌入一个结构体的指针
}

// 定义一个狗的方法 wangwang
func (d *dog) wangwang()  {
	fmt.Printf("%s 在叫：汪汪汪~\n", d.name)
}

func main() {
	var a = dog{
		feet: 4,
		animal: &animal{
			name: "旺财",
		},
	}
	a.wangwang()
	a.move()
}
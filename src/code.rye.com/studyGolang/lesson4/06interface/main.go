package main

import "fmt"

// 为什么要用接口
// Cat结构体
type Cat struct {}

// Say 猫会叫
func (c Cat) Say() string {
	return "喵喵喵"
}

// Dog结构体
type Dog struct{}

// Say 狗也会叫
func (d Dog) Say() string {
	return "汪汪汪"
}

type animal interface {
	Say() string
}

func main()  {
	// c := Cat{}
	// fmt.Println("猫：", c.Say())
	// d := Dog{}
	// fmt.Println("狗：", d.Say())

	var animalList []animal
	c := Cat{} // 造一个猫
	d := Dog{} // 造一个狗

	animalList = append(animalList, c, d)
	fmt.Println(animalList)

	// 
	for _, item := range animalList {
		ret := item.Say()
		fmt.Println(ret)
	}
}
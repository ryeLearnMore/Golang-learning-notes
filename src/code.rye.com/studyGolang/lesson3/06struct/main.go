package main
// 前言：
/*
Go语言中没有类的概念，也不支持类的继承等面向对象的概念。
Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性
*/
import "fmt"

// 结构体
// 创建新的类型要使用type关键字
// 注：结构体字段之间不用逗号分隔
type student struct {
	name string
	age int
	gender string
	hobby []string
}
func main()  {
	var haojie = student{
		name: "豪杰",
		age: 19,
		gender: "男",
		hobby: []string{"篮球", "足球", "双色球"},
	}
	fmt.Println(haojie)

	// 结构体支持.访问属性
	fmt.Println(haojie.hobby)

	// 实例化方法1
	// struct是值类型
	// 如果初始化时没有给属性（字段）设置对应的初始值，那么各个字段的值为默认的0值
	var wangzhe = student{}
	fmt.Println(wangzhe.name)
	fmt.Println(wangzhe.age)
	fmt.Println(wangzhe.hobby)

	// 实例化方法2 new(T) t:表示类型或结构体
	var mayi = new(student) // new关键字取地址
	fmt.Println(mayi) // &{ 0  []}
	// (*mayi).name
	mayi.name = "麻衣"
	mayi.age = 19
	fmt.Println(mayi.name, mayi.age)

	// 实例化方法3
	var nazha = &student{}
	fmt.Println(nazha)
	nazha.name = "娜扎"
	fmt.Println(nazha.name)

	// 结构体初始化
	// 只填值初始化
	var stu1 = student{
		"新垣结衣",
		18,
		"女",
		[]string{"我老婆"},
	}
	fmt.Println(stu1.name, stu1.hobby)
	// 键值对初始化
	var stu2 = &student{
		name: "石原里美",
		age: 18,
	}
	fmt.Println(stu2.name, stu2.age, stu2.gender)
}
package main

// for循环 
import "fmt"

func main()  {
	age := 3
	for age > 0 { // 相当于别的语言中的while循环
		fmt.Println(age)
		age -= 1
	}

	// switch
	switchDemo1()
	switchDemo2()
	switchDemo3()

	// switch test
	testSwitch2()
}

func switchDemo1()  {
	finger := 3

	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

func testSwitch2()  {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("无效输入！")
	}
}

func switchDemo2()  {
	age := 20
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好！")
	}
}

// fallthrough 语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
// 输出:
// a
// b
func switchDemo3() {
	s := "a"
	switch {
	case s == "a": // 成立
		fmt.Println("a")
		fallthrough // 无条件的执行下面case的语句
	case s == "b": 
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
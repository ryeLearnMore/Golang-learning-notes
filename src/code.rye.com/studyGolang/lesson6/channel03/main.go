package main

import "fmt"

// 接收值时判断通道是否关闭

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 不像打开的文件必须在代码中显式的关闭；channel可以被垃圾回收机制回收
}

func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)
	// 方法1
	/*
		// 利用for循环去通道ch1中接收值
		for {
			ret, ok := <-ch1
			if !ok {
				break
			}
			fmt.Println(ret)
		}
	*/

	// 方法2
	for ret := range ch1 {
		fmt.Println(ret)
	}
}

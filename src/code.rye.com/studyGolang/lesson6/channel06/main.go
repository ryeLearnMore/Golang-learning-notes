package main

import "fmt"

// select 练习2

func main() {
	// 声明了一个存放int类型 容量为1的通道
	var ch = make(chan int, 1)
	// for 循环十次
	for i := 0; i < 10; i++ {
		select {
		case ch <- i: // 尝试往ch1中发送数据
		case ret := <-ch: // 尝试从ch1中接收数据
			fmt.Println(ret)
		}
	}
}

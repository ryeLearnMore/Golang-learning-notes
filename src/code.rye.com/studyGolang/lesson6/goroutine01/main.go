package main

import (
	"fmt"
	// "time"
	"sync"
)

// 启动goroutine
// 作用：利用sync.WaitGroup实现优雅的等待
var wg sync.WaitGroup // 是一个结构体，里面有一个计数器

func hello(i int) {
	defer wg.Done() // 相当于计数器-1
	fmt.Println("Hello didi!", i)
	// wg.Done() // 相当于计数器-1
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go hello(i) // 1. 创建一个goroutine 2. 在新的goroutine中执行hello函数
	}
	fmt.Println("Hello main func.")
	// time.Sleep(time.Second)
	// 等hello执行完（执行hello函数的那个goroutine执行完）
	wg.Wait() // 阻塞，一直等待所有的goroutine结束
}

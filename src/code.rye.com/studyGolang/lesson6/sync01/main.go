package main

import (
	"fmt"
	"sync"
)

var x int64 // 定义全局变量x
var wg sync.WaitGroup

// 定义一个互斥锁
var lock sync.Mutex

// 定义一个函数 对全局的变量x做累加的操作
func add() {
	for i := 0; i < 5000; i++ {
		// 操作流程
		/*
			1. 从内存中找到x的值
			2. 执行 +1 操作
			3. 把结果赋值给x写到内存
		*/

		// 上锁后
		lock.Lock() // 上锁
		x++
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x) // 不加锁的话，输出的值有时为10000，有时会小于10000
}

package main

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// func a() {
// 	defer wg.Done()
// 	for i := 1; i < 10; i++ {
// 		fmt.Println("A:", i)
// 	}
// }

func b() {
	defer wg.Done()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("B:", i)
	// }
	var i int64
	for {
		i++
	}
}

func main() {
	// runtime.GOMAXPROCS(1) // 设置我的go程序只用一个逻辑核心，m:n中设置n为1
	runtime.GOMAXPROCS(1) // 设置我的go程序只用一个逻辑核心，m:n中设置n为1，如果不限制，CPU会达到100%
	wg.Add(10)
	// go a()
	for i := 0; i < 10; i++ {
		go b()
	}
	// time.Sleep(time.Second)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup

func read()  {
	defer wg.Done()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 10)
}

func write()  {
	defer wg.Done()
	x ++
	time.Sleep(time.Millisecond * 50)
}

func main()  {
	start := time.Now()
	// 写10次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("耗费了%v.", end.Sub(start))
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁
var rwLock sync.RWMutex // 读写互斥锁：多个goroutine同时读加的是读锁，写的时候加的是写锁

func read()  {
	defer wg.Done()
	// lock.Lock()
	rwLock.RLock() // 加读锁
	fmt.Println(x)
	// lock.Unlock()
	time.Sleep(time.Millisecond * 10)
	rwLock.RUnlock() // 释放读锁
}

func write()  {
	defer wg.Done()
	// lock.Lock()
	rwLock.Lock() // 加写锁
	x++
	// lock.Unlock()
	time.Sleep(time.Millisecond * 50)
	rwLock.Unlock() // 释放写锁
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

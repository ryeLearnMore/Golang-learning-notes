package main

import (
	"fmt"
	"strconv"
	"sync"
)
// sync.Map

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int)  {
	m[key] = value
}

func main()  {
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ { // -->此处值设置的过大，就会报错，因为此时的map并不安全
		wg.Add(1)
		go func(n int) { // 将int类型转换为字符串类型
			key := strconv.Itoa(n)
			set(key, n) // 给map设置键值对
			fmt.Printf("k=:%v, v:=%v\n", key, get(key)) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
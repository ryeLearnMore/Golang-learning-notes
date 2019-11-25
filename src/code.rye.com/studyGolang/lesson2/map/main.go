package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main()  {
	// initMap()
	// judegeIs()
	// forRangeMap()
	// deleteMap()
	// application1()
	// mapSlice()
	// sliceMap()
	countWords()
}

// map的初始化
func initMap()  {
	// 只声明map类型，但是没有初始化，则map初始值就是nil
	var a map[string]int
	fmt.Println(a == nil) // true

	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	// map中添加键值对
	a["娜扎"] = 100
	a["小王子"] = 200
	fmt.Println("a: ", a) // a:  map[小王子:200 娜扎:100]
	fmt.Printf("a:%v\n", a) // a:map[娜扎:100 小王子:200]
	fmt.Printf("a:%#v\n", a) // a:map[string]int{"小王子":200, "娜扎":100}
	fmt.Printf("a:%T\n", a) // a:map[string]int

	// 声明map的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b) // b:map[int]bool{1:true, 2:false}
	fmt.Printf("b:%T", b) // b:map[int]bool

	// 错误写法
	var c map[int]int
	c[100] = 200 // c这个map没有初始化，所以并没有在内存中申请到空间，所以就不能直接操作
	fmt.Println(c)
}

// 判断某个键是否存在
func judegeIs()  {
	var scoreMap = make(map[string]int, 9)
	scoreMap["A"] = 100
	scoreMap["B"] = 200

	// 判断C在不在scoreMap中？
	// value, ok := scoreMap["C"] // 0 false
	value, ok := scoreMap["A"] // 100 true
	fmt.Println(value, ok) 
	if ok {
		fmt.Println("C在scoreMap中", value)
	} else {
		fmt.Println("查无此人！")
	}
}

// for range 遍历方法
func forRangeMap()  {
	var scoreMap = make(map[string]int, 9)
	scoreMap["A"] = 100
	scoreMap["B"] = 200
	// map是无序遍历的，和键值对添加的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 只遍历map中的key
	for k := range scoreMap {
		fmt.Println(k)
	}
	// 之便利map中的value
	for _, v := range scoreMap {
		fmt.Println(v)
	}
}

func deleteMap()  {
	// 删除某个键值对
	var scoreMap = make(map[string]int, 9)
	scoreMap["A"] = 100
	scoreMap["B"] = 200

	delete(scoreMap, "A")
	fmt.Println(scoreMap)
}

// 按照某个固定顺序遍历map
func application1()  {
	var scoreMap = make(map[string]int, 100)

	// 添加50个键值对
	for i := 0; i < 50; i ++ {
		// fmt.Sprintf的用法：http://c.biancheng.net/view/41.html
		key := fmt.Sprintf("stu%02d", i) 
		value := rand.Intn(100) // 0-99的随机整数
		scoreMap[key] = value
	}

	// for key, value := range scoreMap {
	// 	fmt.Println(key, value)
	// }

	// 按照key从大到小的顺序去遍历scoreMap
	// 1. 先取出所有的key存放到切片中
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	// 2. 对key排序
	sort.Strings(keys) // keys排序后成为有序的切片
	// 3. 按照排序后的key对scoreMap排序
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// 元素类型为map的切片
func mapSlice()  {
	var mapSlice = make([]map[string]int, 8, 8) // 切片的初始化
	// 此时 map中的元素为[nil nil nil nil nil nil nil nil]
	// fmt.Println(mapSlice[0] == nil) // true
	// 还需要完成内部map的初始化
	mapSlice[0] = make(map[string]int, 8) // 完成map的初始化
	mapSlice[0]["A"] = 100
	fmt.Println(mapSlice) // [map[A:100] map[] map[] map[] map[] map[] map[] map[]]

}

// 值为切片类型的map
func sliceMap()  {
	 var sliceMap = make(map[string][]int, 8)
	 v, ok := sliceMap["china"]
	 if ok {
		 fmt.Println(v)
	 } else {
		 // slice中没有china这个键
		 sliceMap["china"] = make([]int, 8) // 对切片初始化
		 sliceMap["china"][0] = 100
		 sliceMap["china"][1] = 100
		 sliceMap["china"][3] = 100
	 }

	 // 遍历sliceMap
	 for k, v := range sliceMap {
		 fmt.Println(k, v) // china [100 100 0 100 0 0 0 0]
		 fmt.Printf("%T", sliceMap) // map[string][]int
	 }
}

// 统计一个字符串中每个单词出现的次数
func countWords()  {
	/*
	example: "how do you do"中每个单词出现的次数
	步骤：
	0. 定义一个map[string]int
	1. 字符串中有哪些单词
	2. 遍历单词做统计
	*/
	// 步骤0
	var s = "how do you do"
	var wordCount = make(map[string]int, 8)
	// 步骤1
	words := strings.Split(s, " ")
	// 步骤2
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			// map中有这个单词的统计记录
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}

	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
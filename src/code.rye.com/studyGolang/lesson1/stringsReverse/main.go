package main

// 字符串反转操作
import "fmt"

func main()  {
	s1 := "hello"
	// 法一， 空间复杂度较高，需要申请额外的内存空间
	byteArray := []byte(s1)
	s2 := ""
	for i := len(byteArray) - 1; i >= 0; i -- {
		s2 += string(byteArray[i]) // 102 --> "h"，所以需要强转
	}
	fmt.Println(s2)

	// 法二, 类似于二分法，进行换位
	length := len(s1)
	for i := 0; i < length / 2; i ++ {
		byteArray[i], byteArray[length - i - 1] = byteArray[length - i - 1], byteArray[i]
	}
	fmt.Println(string(byteArray))

}
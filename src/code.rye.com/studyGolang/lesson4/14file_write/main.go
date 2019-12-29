package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.OpenFile("xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	str := "Hello 沙河"
	file.Write([]byte("哎， 张迪也不理我啊\n"))
	file.WriteString(str)
}

package main

import (
	"fmt"
	"encoding/json"
)


// json序列化

// student 学生(必须大写)
// 如果想让输出的信息全为小写，可以：定义元信息：json tag

type Student struct{
	ID int `json:"id"` // 注：`json: "id"`有空格会报错
	Gender string `json:"gender"`
	Name string `json:"name"`
}

func main()  {
	var stu1 = Student{
		ID: 1, 
		Gender: "男",
		Name: "豪杰",
	}

	// 序列化：把编程语言里面的数据转换成JSON格式的字符串
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("Json格式化出错啦！")
		fmt.Println(err)
	}

	fmt.Println(v) // []byte
	fmt.Println(string(v)) // 把[]byte转换成string // {"ID":1,"Gender":"男","Name":"豪杰"}
	fmt.Printf("%#v", string(v)) // "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"豪杰\"}"

	// 反序列化：把满足JSON格式的字符串转换成 当前编程语言里面的对象
	str := "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"豪杰\"}"
	var stu2 = &Student{}
	json.Unmarshal([]byte(str), stu2)
	fmt.Println(stu2) // "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"豪杰\"}"&{1 男 豪杰}
	fmt.Printf("%T\n", stu2) // *main.Student

}
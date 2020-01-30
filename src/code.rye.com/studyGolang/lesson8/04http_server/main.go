package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTP server

func search(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./form.html")
	if err != nil {
		fmt.Println("read html file failed: ", err)
		return
	}
	w.Write(data)
}

func index(w http.ResponseWriter, r *http.Request) {
	// r: 代表跟请求相关的所有内容
	// 获取注册的信息

	// 获取请求的方法
	fmt.Println("请求的方法：", r.Method)
	r.ParseForm() // 解析表格
	// 获取表单中的数据
	fmt.Printf("%#v\n", r.Form)
	usernameValue := r.Form.Get("username")
	pwdValue := r.Form.Get("pwd")
	fmt.Println(r.Form["pwd"])
	fmt.Println(usernameValue, pwdValue)
	// 校验用户输入信息的有效性
	w.Write([]byte("index"))
}

func main() {
	http.HandleFunc("/web", search)
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

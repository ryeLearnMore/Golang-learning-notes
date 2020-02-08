package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTML render

func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "嘿嘿嘿",
	})
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "哈哈哈",
	})
}

func helloHandler(c *gin.Context) {
	// 声明一个结构体类型
	type userinfo struct {
		Name     string `json:"username"`
		Password string `json:"pwd"`
	}
	// 实例化一个userinfo对象
	u1 := userinfo{
		Name:     "didi",
		Password: "123456",
	}

	c.JSON(http.StatusOK, u1)
}

func xmlHandler(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"msg": "xml",
	})
}

func main() {
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	// 设置静态文件的目录
	// 第一个参数是代码里使用的路径，第二个参数是实际保存静态文件的路径
	r.Static("/myWife", "./statics")
	r.GET("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/hello", helloHandler)
	r.GET("/xml", xmlHandler)
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})

	r.Run()
}

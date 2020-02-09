package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryStringHandler(c *gin.Context) {
	// 获取query string参数
	nameVal := c.DefaultQuery("name", "didi") // 查不到就用指定的默认值（第二个参数）
	cityVal := c.Query("city")
	c.JSON(http.StatusOK, gin.H{
		"name": nameVal,
		"city": cityVal,
	})
}

func formHandler(c *gin.Context) {
	// 提取form表单数据
	nameVal := c.PostForm("name")
	cityVal := c.DefaultPostForm("city", "beipiao")
	c.JSON(http.StatusOK, gin.H{
		"name": nameVal,
		"city": cityVal,
	})
}

func paramsHandler(c *gin.Context) {
	// 提取路径参数
	actionVal := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": actionVal,
	})
}

// 参数相关示例
func main() {
	r := gin.Default()
	// query string: http://127.0.0.1:8080/search?name=didi&city=beijing
	r.GET("/query_string", queryStringHandler)
	// form params: HTML页面上form表单提交数据
	r.POST("/form", formHandler)
	// URL参数：/book/list	/book/new	/book/delete
	r.GET("/book/:action", paramsHandler)
	r.Run()
}

package main

// BookManagementSystem
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func bookListHandler(c *gin.Context) {
	/*
		连数据库
		查数据
		返回给浏览器
	*/

	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	// 返回数据
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"data": bookList,
	// })
	c.HTML(http.StatusOK, "book/book_list.tmpl", gin.H{
		"code": 0,
		"data": bookList,
	})
}

func newBookHandler(c *gin.Context) {
	// 给用户返回一个添加书籍页面的处理函数
	c.HTML(http.StatusOK, "book/new_book.html", nil)
}

func createBookHandler(c *gin.Context) {
	// 创建书籍的处理函数
	// 从form表单取数据
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	price, err := strconv.ParseFloat(priceVal, 64)
	if err != nil {
		msg := "无效的价格参数"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	// 拿到数据 插入数据库
	err = insertBook(titleVal, price)
	if err != nil {
		msg := "插入数据失败，请重试！"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	// 插入数据成功
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

// 删除具体某本书
func deleteBookHandler(c *gin.Context) {
	// 取query string 参数
	idStr := c.Query("id")
	idVal, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	// id是个正经数字
	// 去数据库删除具体的记录
	err = deleteBook(idVal)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	// 删除数据成功后，重定向到书籍列表页
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func main() {
	// 程序一启动就应该连接数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	// 查看所有的书籍数据
	r.GET("/book/list", bookListHandler)
	// 返回一个页面给用户填写新增的书籍信息
	r.GET("/book/new", newBookHandler)
	r.POST("/book/new", createBookHandler)
	r.GET("/book/delete", deleteBookHandler)
	r.Run()
}

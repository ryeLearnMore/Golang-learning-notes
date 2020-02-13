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

// day11
func editBookHandler(c *gin.Context) {
	// 1. 获取需要编辑的书的ID值
	idStr := c.Query("id")
	if len(idStr) == 0 {
		// 请求中没有携带需要用的数据，所以该请求无效
		c.String(http.StatusBadRequest, "无效的请求")
		return
	}
	// HTTP 请求传过来的参数通常都是string类型，要根据自己的需要去转换成相应的数据类型
	bookID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// 请求中没有携带需要用的数据，所以该请求无效
		c.String(http.StatusBadRequest, "无效的请求")
		return
	}
	if c.Request.Method == "POST" {
		// 1. 获取用户提交的数据
		titleVal := c.PostForm("title")
		priceStr := c.PostForm("price")

		priceVal, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "无效的价格信息")
			return
		}
		// 2. 去数据库更新对应的书籍数据
		// 难点：id去哪了？需在html中指定一下
		err = editBook(titleVal, priceVal, bookID)
		if err != nil {
			c.String(http.StatusBadRequest, "更新数据失败")
			return
		}

		// 3. 跳转回/book/list页面查看是否修改成功
		// 注： 相同网站跳转，可以写相对路径。不一样的网站跳转：需要些绝对路径
		c.Redirect(http.StatusMovedPermanently, "/book/list")
	} else {
		// 需求：编辑时需要给模板上渲染原来的旧数据
		// // 1. 取到用户编辑的是哪一本书？可以从querystring取到编辑的书的ID值
		// idStr := c.Query("id")
		// if len(idStr) == 0{
		// 	// 请求中没有携带需要用的数据，所以该请求无效
		// 	c.String(http.statusBadRequest, "无效的请求")
		// 	return
		// }
		// // HTTP 请求传过来的参数通常都是string类型，要根据自己的需要去转换成相应的数据类型
		// bookID, err := strconv.ParseInt(idStr, 10, 64)
		// if err != nil {
		// 	// 请求中没有携带需要用的数据，所以该请求无效
		// 	c.String(http.StatusBadRequest, "无效的请求")
		// 	return
		// }
		// 2. 根据id取到书籍的信息
		bookObj, err := queryBookByID(bookID)
		if err != nil {
			c.String(http.StatusBadRequest, "无效的书籍ID")
			return
		}
		// 3. 把书籍数据渲染到页面上
		c.HTML(http.StatusOK, "book/book_edit.html", bookObj)
	}
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
	// day11
	r.Any("/book/edit", editBookHandler)
	r.Run()
}

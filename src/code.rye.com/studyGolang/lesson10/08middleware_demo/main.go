package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/index",
	})
}
func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	fmt.Println(c.MustGet("key").(string))
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/home",
	})
}

func castTime(c *gin.Context) {
	start := time.Now()
	c.Set("key", "value")
	c.Next() // 运行下一个Handler函数

	// 统计耗时
	cast := time.Since(start)
	fmt.Println("cast:", cast)
}

// 路由分组
func main() {
	// r := gin.New()
	r := gin.Default() // 默认使用了2个中间件:1.记日志 2.recover

	// 下面两个GET都会提前计算每个函数执行的时间
	r.Use(castTime)
	// 根据URL分执行的函数
	shoppingGroup := r.Group("/shopping")
	{
		// shoppingGroup.GET("/index", castTime, shopIndexHandler) // 执行单个函数的时间
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}

	r.Run()
}

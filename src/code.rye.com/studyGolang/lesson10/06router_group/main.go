package main

import (
	"github.com/gin-gonic/gin"
)

// 路由分组

// func shopIndexHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "shopping/index",
// 	})
// }

// func shopHomeHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "shopping/home",
// 	})
// }

// func liveIndexHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "live/index",
// 	})
// }

// func liveHomeHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "live/home",
// 	})
// }

func main() {
	r := gin.Default()
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}

	liveGroup := r.Group("/live")
	{
		liveGroup.GET("/index", liveIndexHandler)
		liveGroup.GET("/home", liveHomeHandler)
	}

	// 多层嵌套路由
	v1 := r.Group("/v1")
	{
		v1Shopping := v1.Group("/shopping")
		{
			v1Shopping.GET("/home", shopHomeHandler)
		}
	}
	r.Run()
}

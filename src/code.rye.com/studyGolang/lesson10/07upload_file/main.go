package main

// 文件上传

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	// 提取用户上传的文件
	fileObj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	// fileobj: 上传的文件对象
	// fileobj.Filename // 拿到上传的文件名
	filePath := fmt.Sprintf("./%s", fileObj.Filename)
	// 保存文件到本地的路径
	c.SaveUploadedFile(fileObj, filePath)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})

}

// 文件上传示例
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	r.POST("/upload", uploadHandler)
	r.Run()
}

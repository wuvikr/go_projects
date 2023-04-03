package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	// 加载html文件
	e.LoadHTMLFiles("./upload.html")

	e.GET("/upload", func(ctx *gin.Context) {
		// 返回html页面
		ctx.HTML(http.StatusOK, "upload.html", nil)
	})

	e.POST("/upload", func(ctx *gin.Context) {
		// 获取表单上传的文件
		file, err := ctx.FormFile("f1")
		log.Println(file.Filename)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			dst := path.Join("./", file.Filename)
			// 将上传文件保存到指定目录
			ctx.SaveUploadedFile(file, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
			})
		}
	})

	e.Run(":9090")
}

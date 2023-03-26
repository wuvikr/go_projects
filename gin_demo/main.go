package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createbook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create book!",
	})
}

func main() {
	e := gin.Default() // 返回默认路由引擎

	e.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "get book!",
		})
	})

	e.POST("/book", createbook)

	e.PUT("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update book!",
		})
	})

	e.DELETE("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "delete book!",
		})
	})

	// 启动服务
	e.Run("127.0.0.1:9090")
}

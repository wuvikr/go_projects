package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin获取query string
func main() {
	e := gin.Default()

	e.GET("/web", func(ctx *gin.Context) {
		// http://127.0.0.1:9090/web?query=杨超越
		// http://127.0.0.1:9090/web?query=杨超越&age=22
		// name := ctx.Query("query") // 获取浏览器请求携带的query string
		// name := ctx.DefaultQuery("query", "nobody") // 取不到query的值，就返回默认值
		name, ok := ctx.GetQuery("query") // 第二个返回值为bool类型，取不到query,会返回false
		if !ok {
			name = "nobody"
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	e.Run("127.0.0.1:9090")
}

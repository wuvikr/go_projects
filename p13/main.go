package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取请求的Path（URI）路径
func main() {
	router := gin.Default()

	router.GET("/blog/:year/:month", func(ctx *gin.Context) {
		year := ctx.Param("year") // string类型
		month := ctx.Param("month")
		ctx.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	router.Run("127.0.0.1:9090")
}

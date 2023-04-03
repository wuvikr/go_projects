package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https:www.baidu.com")
	})

	router.Run(":9090")
}

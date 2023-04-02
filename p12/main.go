package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.LoadHTMLFiles("./login.html", "./index.html")

	e.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	e.POST("/login", func(ctx *gin.Context) {
		// 获取form表单提交的数据
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		if username == "小王子" && password == "123456" {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"name": username,
				"pass": password,
			})
		} else {
			ctx.JSON(http.StatusOK, "your username or password is wrong!")
		}

	})

	e.Run("127.0.0.1:9090")
}

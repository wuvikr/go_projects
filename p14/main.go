package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	e := gin.Default()

	e.GET("/query", func(ctx *gin.Context) {
		name := ctx.Query("username")
		pass := ctx.Query("password")

		u := User{
			Username: name,
			Password: pass,
		}

		fmt.Printf("u: %v\n", u)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "query",
		})
	})

	e.POST("/json", func(ctx *gin.Context) {
		var u User

		/*
			shouldBind方法可以基于请求的Content-type，自动识别请求类型，
			并利用反射机制，自动提取请求中的QueryString，Form表单，Json，XML里面的参数到结构体中
		*/
		err := ctx.ShouldBind(&u)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		fmt.Printf("u: %v\n", u)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "json",
		})
	})

	e.POST("/form", func(ctx *gin.Context) {
		var u User

		err := ctx.ShouldBind(&u)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		fmt.Printf("u: %v\n", u)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "form",
		})
	})

	e.Run(":9090")
}

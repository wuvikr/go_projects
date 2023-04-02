package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 构造JSON数据可以使用两种方式
	/*
		方法一： 使用map
		data := map[string]interface{}{
			"name":    "小王子",
			"age":     18,
			"message": "hello !",
		}
	*/
	engine.GET("/json-map", func(ctx *gin.Context) {
		// gin.H是gin框架内置的一个map[string]interface{}类型的类型定义,方便开发者使用
		// type H map[string]any
		ctx.JSON(http.StatusOK, gin.H{
			"name":    "小王子",
			"age":     18,
			"message": "hello",
		})

	})

	// 方法二：使用结构体
	type msg struct {
		Name    string `json:"name"` // 使用tag来做结构体字段定制化
		Age     int
		Message string
	}

	engine.GET("/json-struct", func(ctx *gin.Context) {
		data := msg{
			"张三",
			22,
			"法外狂徒",
		}
		ctx.JSON(http.StatusOK, data)
	})

	engine.Run("127.0.0.1:9090")
}

package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 遇事不决 写注释

// 定义一个User类型
type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1. 解析模板文件
	t, err := template.ParseFiles("./Templates/hello.tmpl")
	if err != nil {
		fmt.Printf("Parse temlate failed, err:%v\n", err)
	}

	// 定义传入模板的数据内容
	u1 := User{
		Name:   "小王子",
		Age:    18,
		Gender: "男",
	}

	m1 := map[string]interface{}{
		"name":   "zhangsan",
		"age":    12,
		"gender": "女",
	}

	hobbyList := []string{
		"唱",
		"跳",
		"rap",
		"篮球",
	}

	// 2. 渲染模板
	err2 := t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err2)
	}
}

func pra(w http.ResponseWriter, r *http.Request) {

	// 定义模板自定义函数
	praise := func(name string) string {
		return name + "你真棒！"
	}

	// 定义模板对象
	t := template.New("f.tpl")

	// 将自定义函数注册到模板对象中
	t.Funcs(template.FuncMap{
		"pra": praise,
	})
	t2, err := t.ParseFiles("./Templates/f.tpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v/n", err)
	}

	var name = "小王子"

	// 渲染模板
	t2.Execute(w, name)

}

func tem(w http.ResponseWriter, r *http.Request) {
	// 1. 解析模板文件
	t, err := template.ParseFiles("./Templates/t.tmpl", "./Templates/ul.tmpl")
	if err != nil {
		fmt.Printf("Parse temlate failed, err:%v\n", err)
	}

	var name = "小王子"

	// 渲染模板
	t.Execute(w, name)
}

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	// 需要先加载根模板
	t, err := template.ParseFiles("./Templates/base.tmpl", "./Templates/index.tmpl")

	if err != nil {
		fmt.Printf("Parse temlate failed, err:%v\n", err)
	}

	var name = "小王子"

	// 渲染模板
	t.ExecuteTemplate(w, "index.tmpl", name)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./Templates/base.tmpl", "./Templates/home.tmpl")

	if err != nil {
		fmt.Printf("Parse temlate failed, err:%v\n", err)
	}

	var name = "小王子"

	// 渲染模板
	t.ExecuteTemplate(w, "home.tmpl", name)
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 创建模板对象并添加自定义方法
	t := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// 解析模板
	t2, err := t.ParseFiles("./Templates/xss.tmpl")

	if err != nil {
		fmt.Printf("Parse temlate failed, err:%v\n", err)
		return
	}

	content1 := "<script>alert(123);</script>"
	content2 := "<a href='https://www.baidu.com'>百度</a>"

	// 渲染模板
	t2.Execute(w, map[string]string{
		"con1": content1,
		"con2": content2,
	})
}

/*
// net/http包实现
func main() {
	http.HandleFunc("/", sayHello)

	http.HandleFunc("/pra", pra)

	http.HandleFunc("/tem", tem)

	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	http.HandleFunc("/xss", xss)

	// 启动服务
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Printf("HTTP serve start failed, err:%v", err)
		return
	}
}
*/

// gin框架实现
func main() {
	r := gin.Default()

	// 添加模板自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// 模板解析
	// r.LoadHTMLFiles("Templates/users/gin_tmp.tmpl", "Templates/users/gin_tmp.tmpl")
	r.LoadHTMLGlob("Templates/**/*") // 使用正则表达式匹配来加载文件

	r.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/gin_tmp.tmpl", gin.H{
			"title": "www.index.com",
		})
	})

	r.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/gin_tmp.tmpl", gin.H{
			"title": "www.posts.com",
		})
	})

	// 运行服务
	r.Run("127.0.0.1:9090")
}

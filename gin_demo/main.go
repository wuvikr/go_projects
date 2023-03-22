package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// /hello Handle的返回内容
func sayHello(w http.ResponseWriter, r *http.Request) {
	byte, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = fmt.Fprintln(w, string(byte))
}

func main() {
	// 创建一个路径为 /hello 的Handle
	http.HandleFunc("/hello", sayHello)

	// 创建一个serve
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Printf("http serve failed, err: %v\n", err)
		return
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	addr := "0.0.0.0:8080"

	//favicon.ico  文件不做处理
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	})

	// http.HandleFunc()   这是一个函数调用， 里面定义了一个匿名函数

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//请求行
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("method", r.Method)
		fmt.Println("url", r.URL)
		fmt.Println("protocol", r.Proto)

		//请求头
		// fmt.Println("Header:", r.Header)  //得到的是一个map
		header := r.Header
		fmt.Println(header.Get("User-Agent")) // 获取具体的值
		fmt.Println(r.Host)

		fmt.Fprint(w, time.Now().Format("2006-01-02 15:06:05"))

		//请求体
		fmt.Println("body:")
		io.Copy(os.Stdout, r.Body) // 获取body 的数据
	})

	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}

/*
http.ReponseWriter  http 实现了Io接口  所以可以使用FPrint

method GET
url /favicon.ico   // 会有一个默认请求， 这是浏览器的行为；是在浏览器左上角的一个图标
protocol HTTP/1.1


curl -XGET "localhost:8080/get"
curl -XPUT "localhost:8080/put"
curl -XPOST "localhost:8080/post"
curl -XDELETE "localhost:8080/delete"
curl -Xxxxxxx "localhost:8080/xxxxx"  // 自定义的方法
curl -I "localhost:8080/head"  // head 方法

*/

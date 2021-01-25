package main

import (
	"fmt"
	"net/http"
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
		fmt.Println(strings.Repeat("-", 30))
		// 1 提交数据方式
		/*
			在URL中传递数据
			url?argname1=argvalue&argname2=argvalue2
		*/
		r.ParseForm()       // 解析参数
		fmt.Println(r.Form) // 在浏览器中访问http://localhost:8080/?x=1&y=2   得到 map[x:[1] y:[2]]
		fmt.Fprintln(w, time.Now().Format("2006-01-02 15:03:04"))

	})

	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}

/*
在URL中传递数据
url?argname1=argvalue&argname2=argvalue2
*/

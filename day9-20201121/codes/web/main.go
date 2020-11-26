package main

import (
	"fmt"
	"net/http"
)

// 处理器函数
func Home(w http.ResponseWriter, r *http.Request) {
	// 用户提交的数据 http内容-> go代码转换 http.Request
	w.Write([]byte("hi"))
}

type Help struct{}

func (h *Help) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "help")
}

func main() {
	addr := ":8888"
	// url => 找处理器函数 => 调用处理器函数(http包)
	// 指定url和处理器关系
	// 处理器函数签名由http包定义
	http.HandleFunc("/home", Home)
	http.Handle("/help", new(Help))

	http.ListenAndServe(addr, nil)

}

/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
1 参数中 提供匹配的字符串，处理函数
2 处理函数 有两个参数 一个是 ResponseWriter *Request
ResponseWriter 接口中有 Write([]byte) (int, error) 方法


服务器端
1 开启监听  http  https 都支持
http  http.ListenAndServe(addr, nil)
https  	http.ListenAndServeTLS

2 函数处理
http.Handle
http.HandleFunc
http.Fileserver

客户端
1 http.Get  Post 等方法
2 http.Do  自定义方法

参考：https://www.cnblogs.com/itbsl/p/12175645.html

*/

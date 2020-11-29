package main

import "net/http"

func main() {

	addr := ":8888"
	//http.Dir 类型转换
	http.Handle("/w1/", http.FileServer(http.Dir("./static/"))) // http.Dir  这是一个类型转换的用法，不是函数

	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("./static/"))))
	// ./static/www/

	http.ListenAndServe(addr, nil)
}

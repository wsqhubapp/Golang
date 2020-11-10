// 匿名结构体

package main

import (
	"fmt"
	"time"
)

func main() {
	// 回顾一下匿名函数
	// var c func()
	// c = func() {

	// }
	// 匿名结构体：在定义变量时将类型指定为结构体的结构
	// 1 用于定义  2 web开发 传输变量给前端
	// 定义命名结构体的变量user
	// var user struct {
	// 	id       int
	// 	name     string
	// 	tel      string
	// 	addr     string
	// 	birthday time.Time
	// }

	// fmt.Printf("%T, %#v\n", user, user)
	// //struct { id int; name string; tel string; addr string; birthday time.Time }, struct { id int; name string; tel string; addr string; birthday time.Time }{id:0, name:"", tel:"", addr:"", birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}

	// 初始化
	// 零值
	// 字面量  和标准函数体一样
	//定义匿名结构体变量user
	user := struct {
		id       int
		name     string
		tel      string
		addr     string
		birthday time.Time
	}{10, "kk", "xxxxxxx", "xxxx", time.Now()} //结构体类型{}

	fmt.Printf("%T, %#v\n", user, user)
	//struct { id int; name string; tel string; addr string; birthday time.Time }, struct { id int; name string; tel string; addr string; birthday time.Time }{id:10, name:"kk", tel:"xxxxxxx", addr:"xxxx", birthday:time.Time{wall:0xbfe29f8e0a903b98, ext:997301, loc:(*time.Location)(0x58abc0)}}
}

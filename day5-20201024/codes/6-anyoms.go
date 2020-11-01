// 匿名结构体

package main

import (
	"fmt"
	"time"
)

func main() {

	// var c func()
	// c = func() {

	// }
	// 匿名结构体
	// 1 用于定义  2 web开发 传输变量给前端
	// 定义匿名结构体的变量user
	var user struct {
		id       int
		name     string
		tel      string
		addr     string
		birthday time.Time
	}

	fmt.Printf("%T, %#v\n", user, user)

	// 初始化
	// 零值
	// 字面量  和标准函数体一样

	user = struct {
		id       int
		name     string
		tel      string
		addr     string
		birthday time.Time
	}{10, "kk", "xxxxxxx", "xxxx", time.Now()} //结构体类型{}

	fmt.Printf("%T, %#v\n", user, user)

}

// 结构体 匿名组合
package main

import (
	"time"
)

type Addr struct {
	province string
	street   string
	no       string
}

type Tel struct {
	prifix string
	number string
}

// 匿名组合方式
type User struct {
	id   int
	name string
	Addr // 只指定类型  ##会定义属性名Addr
	Tel
	birthday time.Time
}

func main() {
	// var user User

	// fmt.Printf("%T, %#v\n", user, user) // 会显示类型名 Addr Tel
	// //main.User, main.User{id:0, name:"", Addr:main.Addr{province:"", street:"", no:""}, Tel:main.Tel{prifix:"", number:""}, birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}

	// user.Addr.province = "北京市"
	// fmt.Println(user.Addr.province) //北京市
	// fmt.Println(user.province)      //北京市
	// // 这是匿名函数的优势 可以简写，赋值、访问都是一样的
	// // 这里会涉及到冲突的问题，优先找本地的， 如果两个子结构体的属性都有，程序就报错；那就需要写标准路径

}

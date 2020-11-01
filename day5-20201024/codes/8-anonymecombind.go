// 匿名组合
//
package main

import (
	"fmt"
	"time"
)

type Addr struct {
	province string
	street   string
	no       string
}

//组合的实现

type Tel struct {
	prifix string
	number string
}

// 匿名组合方式
type User struct {
	id   int
	name string
	Addr // 只指定类型  、、会定义属性名
	Tel
	birthday time.Time
}

func main() {
	var user User

	fmt.Printf("%T, %#v\n", user, user) // 会显示类型名 Addr Tel

	user.Addr.province = "北京市"
	fmt.Println(user.Addr.province)
	fmt.Println(user.province) // 这是匿名函数的优势 可以简写，赋值、访问都是一样的
	// 这里会涉及到冲突的问题，优先找本地的， 如果两个子的都有，程序就报错；那就需要写标准路径

	// var user User = User{addr: Addr{province: "山东省"}}
	// fmt.Printf("%T, %#v\n", user, user)

	// fmt.Printf("%T, %#v\n", user, user)

}

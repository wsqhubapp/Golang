// 匿名组合的指针
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
	id    int
	name  string
	*Addr // 只指定类型  、、会定义属性名
	*Tel
	birthday time.Time
}

func main() {
	// var user User

	// fmt.Printf("%T, %#v\n", user, user)

	var user User = User{Addr: &Addr{province: "山东省"}}
	fmt.Printf("%T, %#v\n", user, user)

	fmt.Println(user.province)
	user.province = "beijingshi"
	fmt.Println(user.province)

}

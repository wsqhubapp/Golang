// 组合 命名组合
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

// 定义结构体
type User struct {
	id       int
	name     string
	addr     Addr
	ter      Tel
	birthday time.Time
}

func main() {
	var user User = User{addr: Addr{province: "山东省"}}
	fmt.Printf("%T, %#v\n", user, user)

	fmt.Println(user.addr.province)

	user.addr.province = "北京市"

	fmt.Printf("%T, %#v\n", user, user)

}

// 结构体的组合 命名组合
package main

import (
	"fmt"
	"time"
)

//实际地址是省、市、区、街道、编号等，下面随便定义了几个
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
	//main.User, main.User{id:0, name:"", addr:main.Addr{province:"山东省", street:"", no:""}, ter:main.Tel{prifix:"", number:""}, birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}

	fmt.Println(user.addr.province) //山东省  ##属性的访问

	user.addr.province = "北京市" //属性的赋值

	fmt.Printf("%T, %#v\n", user, user)
	//main.User, main.User{id:0, name:"", addr:main.Addr{province:"北京市", street:"", no:""}, ter:main.Tel{prifix:"", number:""}, birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}

}

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
	*Addr // 只指定类型
	*Tel
	birthday time.Time
}

func main() {
	var user User = User{Addr: &Addr{province: "山东省"}}
	fmt.Printf("%T, %#v\n", user, user)
	//main.User, main.User{id:0, name:"", Addr:(*main.Addr)(0xc00006c330), Tel:(*main.Tel)(nil), birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}
	// 关注这种写法  和显示的结果
	fmt.Println(user.province) // 山东省
	user.province = "beijingshi"
	fmt.Println(user.province) //beijingshi

}

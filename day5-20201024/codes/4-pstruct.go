// 4 结构体指针
package main

import (
	"fmt"
	"time"
)

// 定义结构体
type User struct {
	id       int
	name     string
	addr     string
	ter      string
	birthday time.Time
}

type Addr struct {
}

//new 函数
func NewUser(id int, name, addr, tel string, birthday time.Time) *User {
	return &User{id, name, addr, tel, birthday}
}

func main() {
	// var user *User
	// fmt.Printf("%T, %#v\n", user, user) //*main.User, (*main.User)(nil)

	// // var a int
	// // var p *int

	// var user *User = &User{}            // 零值
	// fmt.Printf("%T, %#v\n", user, user) //*main.User, (*main.User)(nil)

	// var user *User = &User{id: 10, name: "kk"} // 赋值
	// fmt.Printf("%T, %#v\n", user, user)
	// //*main.User, &main.User{id:10, name:"kk", addr:"", ter:"", birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}

	// // new 函数初始化
	// user = new(User)
	// fmt.Printf("%T, %#v\n", user, user)
	// //*main.User, &main.User{id:0, name:"", addr:"", ter:"", birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}
	// // 36 41开头的两段对比 是两种不同初始化的方法 ，详细参考day02 中的指针的代码

	// u2 := new(User)
	// fmt.Printf("%T, %#v\n", u2, u2)

	// // 属性访问
	// fmt.Println(u2.name) //这种方式可以访问的。只是本身没有值而已

	// 属性赋值的格式
	// u2.name = "kk"
	// u2.id = 10

	u3 := NewUser(1, "kk", "", "", time.Now()) // 这个是调用的上面的函数实现的
	fmt.Printf("%T, %#v\n", u3, u3)
	//main.User, &main.User{id:1, name:"kk", addr:"", ter:"", birthday:time.Time{wall:0xbfe29d31abb533f0, ext:1994801, loc:(*time.Location)(0x58abc0)}}
}

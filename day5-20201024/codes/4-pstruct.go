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

func main() {
	// var user *User
	// fmt.Printf("%T, %#v\n", user, user) //*main.User, (*main.User)(nil)

	// // var a int
	// // var p *int

	// var user *User = &User{}            // 零值
	// fmt.Printf("%T, %#v\n", user, user) //*main.User, (*main.User)(nil)

	var user *User = &User{id: 10, name: "kk"} // 赋值
	fmt.Printf("%T, %#v\n", user, user)        //*main.User, (*main.User)(nil)

	// new 函数初始化
	user = new(User)
	fmt.Printf("%T, %#v\n", user, user)

	u2 := new(User)
	fmt.Printf("%T, %#v\n", u2, u2)

	// 属性访问
	fmt.Print(u2.name)

	// u2.name = "kk"
	// u2.id = 10


	u3 := NewUser{1, "kk". "". "", time.Now()}
	fmt.Printf("%T, %#v\n", u3, u3)



}

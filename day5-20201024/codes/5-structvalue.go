// 结构体的值
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
	user := User{} //使用结构体零值初始化结构体值对象
	u2 := user
	u2.name = "kk" // user和u2 是两块不同的内存空间

	fmt.Printf("%#v\n", user)
	//main.User{id:0, name:"", addr:"", ter:"", birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}
	fmt.Printf("%#v\n", u2)
	//main.User{id:0, name:"kk", addr:"", ter:"", birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}

}

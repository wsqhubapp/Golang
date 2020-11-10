// 3 结构体

package main

import (
	"fmt"
	"time"
)

// id  => int
// name => string
// addr => string
// tel => string
// birthday => time.Time
// 结构中有不同的类型 那怎么定义呢？

// 定义结构体
type User struct {
	id       int
	name     string
	addr     string
	ter      string
	birthday time.Time
}

/*
初始化
1) 使用结构体零值初始化结构体值对象
2) 使用结构体字面量初始化结构体值对象

3) 使用 new 函数进行初始化结构体指针对象
4) 使用结构体字面量初始化结构体指针对象
3) 4) 的内容在 4-pstruct.go 中体现
*/

func main() {
	var user User
	fmt.Printf("%T, %v\n", user, user) //main.User, {0    {0 0 <nil>}}  零值是各元素的零值组成的一个结构体的变量

	// 字面量  // 按照属性顺序定义的字面量  必须严格按照结构体属性定义顺序指定，每个属性都必须指定
	user = User{10, "kk", "xian", "xxx", time.Now()}
	fmt.Printf("%T, %v\n", user, user) //main.User, {10 kk xian xxx {13825228249564272832 9974301 0x58abc0}}

	// 按照属性定义的字面量
	user = User{name: "wang", id: 11}
	fmt.Printf("%T, %v\n", user, user)

	// 属性访问和修改
	fmt.Println(user.id) // 访问   11

	user.id = 100             //赋值
	user.name = "wzp"         //赋值
	fmt.Printf("%#v\n", user) //main.User{id:100, name:"wzp", addr:"", ter:"", birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}}

}

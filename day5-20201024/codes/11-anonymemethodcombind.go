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

func (addr *Addr) SetProvince(province string) {
	addr.province = province
}

type Tel struct {
	prefix string
	number string
	// province string
}

/*
如果下面的方法生效的话，user.SetProvince("陕西省") 会报错的。
因为 User 结构体下面 Tel Addr 子结构体中都有province 这个属性，所以user.SetProvince("陕西省")  中间不知道找哪一个
*/
// func (tel *Tel) SetProvince(name string) {
// 	fmt.Println("tel.setprovince")
// }

// 命名组合方式
type User struct {
	id   int
	name string
	Addr
	Tel
	birthday time.Time
	province string
}

/*
写这个目的，是为了对比的。
当这个代码生效的时候， user.SetProvince("陕西省") 调用的是结构体为User方法，达不到修改子结构体（Addr）的目的；
如果一定要达到修改子结构体Addr的话，就需要写全路径 user.Addr.SetProvince("陕西省")  就是调用addr 接收者的方法
*/
// func (user *User) SetProvince(name string) {
// 	fmt.Println("user.setprovince")
// }

func main() {
	var user User
	fmt.Printf("%#v\n", user)
	// user.Addr.SetProvince("陕西省")   //标准的命名组合时候的写法，这种写法不报错
	user.SetProvince("陕西省") // 匿名的命名组合可以这么写；匿名的结构体中间可以省略
	fmt.Printf("%#v\n", user)
}

// 通过10 11 两个go文件来对比 结构体的组合、匿名结构体组合调用方法的时候的两种不同的写法。

/*
嵌套的总结
1  User结构体；Addr结构体；Addr为接收者的方法 此时user.Addr.SetProvince("陕西省") user.SetProvince("陕西省")  都能修改addr的属性
2  User结构体；Addr结构体；Addr为接收者的方法 User为接收者的方法 此时user.SetProvince("陕西省")修改是user的属性；user.Addr.SetProvince("陕西省") 修改addr的属性
3  User结构体；Addr结构体；Tel结构体；Addr为接收者的方法；Tel 为接收者方法 此时user.SetProvince("陕西省") 报错；user.Addr.SetProvince("陕西省") 修改addr的属性
4  User结构体(有province属性)；Addr结构体；Addr为接收者的方法 此时user.Addr.SetProvince("陕西省") user.SetProvince("陕西省") 修改的都是addr的属性



*/

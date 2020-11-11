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
	prefix   string
	number   string
	province string
}

// 命名组合方式
type User struct {
	id       int
	name     string
	addr     Addr
	tel      Tel
	birthday time.Time
	province string
}

func main() {
	var user User
	fmt.Printf("%#v\n", user)
	user.addr.SetProvince("陕西省")
	fmt.Printf("%#v\n", user)
}

/*
main.User{id:0, name:"", addr:main.Addr{province:"", street:"", no:""}, tel:main.Tel{prefix:"", number:"", province:""}, birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}, province:""}
main.User{id:0, name:"", addr:main.Addr{province:"陕西省", street:"", no:""}, tel:main.Tel{prefix:"", number:"", province:""}, birthday:time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}, province:""}
*/

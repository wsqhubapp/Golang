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
	user := User{}
	u2 := user
	u2.name = "kk" // user和u2 是两块不同的内存空间

	fmt.Printf("%#v\n", user)
	fmt.Printf("%#v\n", u2)

}

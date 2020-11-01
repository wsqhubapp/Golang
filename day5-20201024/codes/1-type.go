package main

import "fmt"

// 现有数据类型

type Counter int

// 作用是 1  见类型知含义   2 可以添加方法

type User map[string]string

type Callbakck func()

func main() {
	var name Counter
	fmt.Printf("%T, %v\n", name, name) //main.Counter, 0

	name += 10
	fmt.Printf("%T, %v\n", name, name) //main.Counter, 10

	var num int = 10
	c2 := Counter(num) + name      // 类型转换
	fmt.Printf("%T, %v\n", c2, c2) //main.Counter, 20

	// var user User

	var user User = make(User) //初始化

	fmt.Printf("%T, %v\n", user, user) //main.User, map[]

	calllback := map[string]Callback{}

	callbacks

	var c22 Counter2
	fmt.printf("%T, %v\n", c22, c22)

}

// 1 自定义类型
package main

import "fmt"

// 现有数据类型

type Counter int

// 作用是 1  见类型知含义   2 可以添加方法

type User map[string]string

type Callback func() error

type Counter2 Counter

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

	fmt.Printf("%T, %#v\n", user, user) //main.User, main.User{}  # 空的main.User{}
	user["id"] = "1"                    //赋值
	fmt.Println(user)                   //map[id:1]   打印出值

	callbacks := map[string]Callback{} // 声明一个callbacks
	callbacks["add"] = func() error {  // 赋值为一个函数
		fmt.Println("add")
		return nil
	}
	callbacks["add"]() //add 函数调用

	var c22 Counter2                 //声明一个变量  只是类型是一个自定义类型
	fmt.Printf("%T, %v\n", c22, c22) //main.Counter2, 0

}

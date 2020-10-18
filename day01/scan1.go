package main

import "fmt"

func main() {
	name := ""
	fmt.Print("请输入你的名字：")

	fmt.Scan(&name)

	fmt.Println("你输入的名字： ", name)

	age := 0 // 接收int的话  1）输入数字，2）类型转换
	fmt.Print("请输入你的年龄：")

	fmt.Scan(&age)

	fmt.Println("你输入的年龄： ", age)

	msg := "" // 接收int的话  1）输入数字，2）类型转换
	fmt.Print("请输入你的msg：")

	fmt.Scan(&msg)

	fmt.Println("你输入的msg： ", msg)
}

package main

// 常量

import "fmt"

const statusNew int = 1
const statusDeleted int = 2

// 定义常量（需要初始化值）

func main() {
	const (
		Monday  = 10
		Tuesday // 10    // 如果在一个小括号内，如果是赋值，则使用最近的一个已赋值的常量对应的值进行初始化; 这是变量和常量的区别
	)

	fmt.Println
}

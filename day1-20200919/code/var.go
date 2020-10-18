package main

import "fmt"

var name string = "kk"

/*
常见的4种类型

*/

func main() {
	var (
		age    int = 31
		weight     = 132
	)

	// 短声明 只能是在函数内，可以定义多个
	// height := 168

	// var msg string = "信息"
	var msg string = 信息 //

	fmt.Println(age, weight, msg)
}

//局部变量 声明了 必须使用，否则会报错
//

package main

import "fmt"

func main() {
	// 函数内（块）定义的变量必须使用，否则的话  会报错的
	/*var havename string = "kk"
	var  msg = "hello world"
	var desc string
	*/
	var (
		havename string = "kk"
		msg             = "hello world"
		desc     string
	)

	/*
		x:= "x"
		y:= "y"
	*/

	x, y := "x", "y"

	fmt.Println(havename, msg, desc, x, y)

}

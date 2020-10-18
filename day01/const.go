package main

import "fmt"

// 包级别的常量
const (
	packageName string = "package const"
	packageMsg         = "package"
)

func main() {
	const name string = "kk"
	const msg = "msg" //常量可以不适用
	fmt.Println(name)
	// name = "slience"   /常量不能修改
}

/*
常量可以定义
*/

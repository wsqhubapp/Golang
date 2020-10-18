package main

import "fmt"

func main() {
	var havename string = "kk" //定义了类型并且初始化了值
	fmt.Println("0", havename)

	var zeroString string // 定义变量类型，但是不初始化值，初始化使用类型对应的零值（空字符串 “”）
	fmt.Println("1", zeroString)

	var typeString = "kk" // 定义变量省略类型，不能省略初始化值，通过对应的值类型推到变量的类型
	fmt.Println("2", typeString)

	shortString := "kk" // 短声明，通过对应的值类型推到变量的类型， 必须在函数内，或函数内子块使用，不能在包级别使用
	fmt.Println("3", shortString)
}

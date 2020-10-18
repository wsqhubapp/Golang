package main

import "fmt"

func main() {
	// const B  不允许
	const (
		A = "test"
		B // 使用前一个常量的值进行初始化（B）
		C // 使用前一个常量的值进行初始化（c=>B）  => 表示最终值得意思
		D = "testD"
		E // 使用前一个常量的值进行初始化（E）
		F // 使用前一个常量的值进行初始化（F=>）
	)
	fmt.Println("1", B, C)
	fmt.Println("2", E, F)
}

//  这个就是一种语法设置

package main

import "fmt"

func main() {
	// 学号到名字的映射

	var names map[string]string //nil

	fmt.Printf("%T, %#v\n", names, names)

	// 访问元素
	names["Go3005"]

}

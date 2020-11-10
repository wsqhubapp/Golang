//函数作为一种类型存在
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("%T\n", add) // func(int, int) int

	f := add // 本质是 var f func(int, int) int = add   函数类型 func(int, int) int 在这里整体充当一个类型

	fmt.Printf("%T\n", f) // func(int, int) int
	fmt.Println(f(3, 4))  // 7
}

//
package main

import "fmt"

func main() {
	// 定义函数类型变量，并使用零值nil 进行初始化
	var callback func(n1, n2 int) (r1, r2, r3, r4 int)
	fmt.Printf("%T, %v", callback, callback)

	// callback = calc // 赋值为函数calc
	fmt.Println(callback(5, 2))
	//  ????? 无法使用

}

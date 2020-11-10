// 函数参数的作用域
package main

import "fmt"

func main() {
	name := "kk"
	nums := []int{}
	func() {
		//1.
		fmt.Println(name, nums) // kk, []
		name = "silence"
		nums = []int{1, 2, 3}
		//2.
		fmt.Println(name, nums) // silence, [1, 2, 3]
	}()
	//3.
	fmt.Println(name, nums) // silence, [1, 2, 3]

}

// 作用域 1 func()里面没有值，到上一层去找；2 本地有值，就直接去本地的值 值找上一层，赋值给上一层；3 在第二步已经赋值了，所以值如注释所示

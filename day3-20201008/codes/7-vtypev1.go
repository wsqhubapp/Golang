package main

import "fmt"

func main() {
	name := "kk"
	nums := []int{1, 2, 3}

	func(pname string, pnums []int) {
		// 1.
		fmt.Println(pname, pnums) // kk, [1, 2, 3]
		pname = "silence"
		pnums = []int{1, 2} // 新生成一个切片
		// 2.
		fmt.Println(pname, pnums) // silence, [1, 2]
	}(name, nums)
	/*
		pname := name
		pnums := nums
		pname = "silence"
		pnums = []int{1, 2}
	*/

	// 3.
	fmt.Println(name, nums) // kk [1 2 3]
}

// 从作用域的角度理解，就没有问题

package main

import "fmt"

func main() {
	name := "kk"
	nums := []int{1, 2, 3} // 切片

	func(pname string, pnums []int) {
		// 1.
		fmt.Println(pname, pnums) // kk, [1, 2, 3]
		pname = "silence"
		pnums[0] = 100
		// 2.
		fmt.Println(pname, pnums) // silence [100, 2, 3]
	}(name, nums)
	/*
		pname := name
		pname = "silence"

		pnums := nums
		pnums[0] = 100
	*/

	// 3.
	fmt.Println(name, nums) // kk, [100, 2, 3]
}

// 上面的函数本质是一个 func(){}()  函数调用； 但是值的变化，为什么是这样的 有待弄清楚

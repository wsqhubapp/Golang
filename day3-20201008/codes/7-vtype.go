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
		下面是一个赋值过程
			pname := name
			pnums := nums
			pname = "silence"
			pnums[0] = 100   //数组的赋值的概念，底层用的是同一数组
	*/

	// 3.
	fmt.Println(name, nums) // kk, [100, 2, 3]   # 1的值变为100 是因为在2的时候修改的值，是取的外层的地址
}

// 上面的函数本质是一个 func(){}()  函数调用； 但是值的变化，这就是一个作用域的问题。好理解

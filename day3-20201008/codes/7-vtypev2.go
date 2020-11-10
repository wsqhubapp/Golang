package main

import "fmt"

func main() {
	name := "kk"
	nums := []int{1, 2, 3}

	func(name string, nums []int) {
		// 1.
		fmt.Println(name, nums) // kk [1 2 3]
		name = "silence"
		nums = []int{1, 2} //这个是内部新生成的一个变量，你也可以理解为是vtypev1的pnums ，所以内部这个怎么变和外面的没有关系
		// 2.
		fmt.Println(name, nums) //  silence [1 2 ]
	}(name, nums)
	/*
	   形参name := name
	   形参 nums := nums
	   name = "silence"
	   nums = []int{1,2}
	*/
	// 3.
	fmt.Println(name, nums) //kk [1 2 3]
}

// 这个对比 vtypev1 理解

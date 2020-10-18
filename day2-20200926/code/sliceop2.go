package main

import "fmt"

func main() {

<<<<<<< HEAD
	nums := []int{1, 2, 3, 4, 5}
	nums2 := nums[1:3]

	fmt.Println(nums2)
=======
	nums := []int{1, 2, 3, 4, 5} // len = 5 cap = 5
	nums2 := nums[1:3]           // 切片操作不能超过容量

	fmt.Println(nums2) // [2 3]
>>>>>>> add

}

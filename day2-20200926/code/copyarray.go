package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	nums2 := make([]int, 0, len(nums))
	// 空白标识符
	for _, v := range nums {
		nums2 = append(nums2, v)
	}
	fmt.Println(nums2)

}

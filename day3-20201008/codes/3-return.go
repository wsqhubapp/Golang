// 正常的返回值练习

package main

import "fmt"

func test(stop bool) {
	fmt.Println("A")
	if stop {
		fmt.Println("B")
		return //  结束当前函数
	}

	fmt.Println("C")
}

func main() {
	test(true) // A B   表示 if stop 为真
	fmt.Println("----------")
	test(false) // A C   表示 if stop 为假
}

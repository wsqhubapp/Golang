// 实现go的枚举类型   go本身没有枚举类型

package main

import "fmt"

func main() {
	const (
		A = iota // 在常量组内使用iota 初始化0，每次调用 +1
		B
		C
		D
	)
	fmt.Println(A, B, C, D) // 0 1 2 3

	const (
		Mon = iota
		Tuesd
		Wed
		Thur
		Fir
	)
	fmt.Println(Mon, Tuesd, Wed, Thur, Fir) // 0 1 2 3 4
}

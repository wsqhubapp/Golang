// 枚举值
package main

import "fmt"

func main() {
	const (
		Monday  = iota
		Tuesday = iota
		Sunday  = iota
	)

	fmt.Println("1", Monday, Tuesday, Sunday)

	const (
		Monday2 = iota
		Tuesday2
		Sunday2
	)

	fmt.Println("2", Monday2, Tuesday2, Sunday2)

	const (
		Monday3 = iota + 10
		Tuesday3
		Sunday3
	)

	fmt.Println("3", Monday3, Tuesday3, Sunday3)

	const (
		Monday4 = iota * 10
		Tuesday4
		Sunday4
	)

	fmt.Println("4", Monday4, Tuesday4, Sunday4)

}

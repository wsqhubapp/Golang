// 别名方式  用得比较少
package main

import "fmt"

type Counter = int // 别名

func main() {

	var counter Counter

	fmt.Printf("%T, %v\n", counter, counter) //int, 0

	var (
		r rune
		b byte
	)

	fmt.Printf("%T, %T\n", r, b) //int32, uint8

}

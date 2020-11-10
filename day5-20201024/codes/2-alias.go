//2  别名方式  用得比较少； 这个是自定义类型引申出来的一个点
package main

import "fmt"

type Counter = int // 别名

func main() {

	var counter Counter

	fmt.Printf("%T, %v\n", counter, counter) //int, 0
	var num int = 10
	fmt.Println(counter + num) //10  说明int 和Counter是一个类型，可以进行数字的加法运算

	var (
		r rune
		b byte
	)

	fmt.Printf("%T, %T\n", r, b) //int32, uint8

}

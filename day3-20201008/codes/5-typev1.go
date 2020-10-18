//函数类型
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func main() {
	var f func(int, int) int = add //定义一个变量f  类型是函数 值是add
	var fs []func(int, int) int    // 定义一个变量fs 类型是切片中的函数， 一般不会这么用，一般会自定义类型。 后续使用追踪分析一下

	fs = append(fs, add, mul)

	c := f(2, 3)
	fmt.Println(c) //5

	for _, v := range fs {
		fmt.Println(v(2, 3)) // 5 6
	}

}

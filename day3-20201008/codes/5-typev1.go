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
	var fs []func(int, int) int    // 定义一个变量fs 类型是切片(但是类型是函数)， 一般不会这么用，一般会自定义类型。 后续使用追踪分析一下

	fs = append(fs, add, mul) //这个是切片的一种写法  给fs 这个切片添加两个元素

	c := f(2, 3)
	fmt.Println(c) //5

	for _, v := range fs { // fs 是一个切片
		fmt.Printf("%T, %#v\n", v, v)
		fmt.Println(v(2, 3)) // 5 6  # 5是函数add的结构，6是mul的结果
	}

}

/*
上面轮询的结果
func(int, int) int, (func(int, int) int)(0x4a0260)
5
func(int, int) int, (func(int, int) int)(0x4a0280)
6
*/

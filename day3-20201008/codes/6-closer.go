/*
匿名函数又叫闭包， 是指在函数内定义的匿名函数引用外部函数的变量，只要匿名函数继续使用则外部函数赋值的变量不被自动销毁
*/
package main

import "fmt"

// 生成一个函数
// add(int) int
// left + base
// 变量的生命周期 => 内存中存在的时间
// 定义闭包函数，返回一个匿名函数用于计算与base 元素的和
func addBase(base int) func(int) int {
	return func(num int) int {
		return base + num
	}
}

func main() {
	add2 := addBase(2)       //使用闭包函数
	fmt.Printf("%T\n", add2) //func(int) int
	fmt.Println(add2(10))    //12

	base2 := addBase(10)
	fmt.Println(base2(1), base2(10)) // 11 12  闭包使用
}

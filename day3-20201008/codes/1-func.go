// 函数的定义& 调用
package main

import "fmt"

// 无参&无返回值  定义Hello
func sayHello() {
	fmt.Println("hello")
}

// 有参数& 无返回值  定义hi
func sayHi(name string) {
	fmt.Println("hi, ", name)
}

// 有参数& 有返回值 add
func add(a int, b int) int {
	fmt.Println(a, b)
	return a + b // return 关键字用来向函数调用者返回结果
}

func main() {
	// 调用函数  函数名称（参数[实参]）
	sayHello() // 注意小括号  表示函数的调用
	sayHello()

	sayHi("kk")
	name := "王清"
	sayHi(name)

	c := add(3, 4) //实参  ==》 形参
	fmt.Println(c)

}

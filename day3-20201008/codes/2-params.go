// 函数的参数
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

// 参数类型合并
// 参数中多个连续的参数类型相同
// 只保留最后一个类型，前面连续相同的数据类型可省略
// var a int, b int =》 var a, b int
func addV2(a, b int) int {
	return a + b
}

func test(a int, b string, c int) {

}

// testV2(a string, b string, c int)
func testV2(a, b string, c int) {

}

// fmt.Println
// append
// 1, 2, 3, 4, 5, ....
// 可变参数
/*
某些情况下，函数需要处理形参数量可变，需要运算符... 声明可以变参数函数或在调用传递可变参数
1 定义
定义多个可变参数  不能（只能有一个， 必须定义在形参最后）， 比如 ...int   ...string
可变参数则被初始化为对应类型的切片
2 传递
*/

func addAll(a, b int, args ...int) int {
	fmt.Println(a, b, args)
	fmt.Printf("%T\n", args) // 打印出args的类型  是切片  []int
	// args args 切片
	print(args...) // 打印出切片
	sum := a + b
	for _, v := range args { // 所有的参数相加
		sum += v
	}
	return sum
}

func print(args ...int) { //逻辑是打印出切片 索引编号  索引值
	for i, v := range args {
		fmt.Println(i, v)
	}
}

func main() {
	fmt.Println(addV2(3, 4))                       // 7
	fmt.Println(addAll(1, 2))                      // args 是[]
	fmt.Println(addAll(1, 2, 3))                   // args 是[3]
	fmt.Println(addAll(1, 2, 3, 4))                // args 是[3,4]
	fmt.Println(addAll(1, 2, 3, 4, 5, 6, 7, 8, 9)) // args是 [3, 4, 5, 6, 7, 8, 9]

	nums := []int{2, 3, 4, 5}
	fmt.Println(addAll(1, 2, nums...))

	print(nums...)

}

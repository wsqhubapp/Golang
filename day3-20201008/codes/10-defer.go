/*
defer 关键字用来申明函数，不论函数是否发生错误都在函数执行最后执行（return之前）。
如果defer 声明多个函数，则按照声明的顺序，先声明后执行（堆）
常用来做资源释放，记录日志等工作
*/

package main

import "fmt"

func main() {
	fmt.Println("start")

	// defer 函数调用
	// defer 延迟执行
	// 在函数退出之前执行
	defer func() {
		fmt.Println("defer")
	}()

	fmt.Println("end")
	// start -> end -> defer
}

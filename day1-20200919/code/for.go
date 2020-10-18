package main

import "fmt"

func main() {
	// 要实现  1 + 2 + 3 + 4 ... 100等于多少   现有的知识点，就学过 += 这种赋值运算符，顾实现代码如下
	sum := 0
	for i := 0; i < 101; i++ {
		sum += i // sum = sum + i
	}
	fmt.Println(sum)
}

/*
基本的for 语句

*/

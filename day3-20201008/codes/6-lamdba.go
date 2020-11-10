// 匿名函数
package main

import (
	"fmt"
	"strings"
)

//匿名函数作为print的参数
func print(formatter func(string) string, args ...string) {
	for i, v := range args {
		fmt.Println(i, formatter(v))
	}
}

func add(a, b int) int {
	return a + b
}

func main() {

	// 匿名函数=> 没有名字的函数
	//定义匿名函数并赋值给c
	// c := func() {
	// 	fmt.Println("我是匿名函数")
	// }
	/*
		func c() {
			fmt.Println("我是匿名函数")
		}
	*/

	// fmt.Printf("%T\n", c) //func()
	// c()                   // 我是匿名函数  #调用匿名函数
	// c()                   // 我是匿名函数
	// c()                   // 我是匿名函数

	names := []string{"赵昌建", "kk", "17-赵"}

	star := func(txt string) string {
		return "*" + txt + "*"
	}

	print(star, names...)
	/*
		0 *赵昌建*
		1 *kk*
		2 *17-赵*
	*/

	// 1 + 2
	a, b := 1, 2
	fmt.Println(add(a, b)) //3 a是1  b是2
	fmt.Println(add(1, 2)) //3

	print(func(txt string) string {
		return "|" + txt + "|"
	}, names...)
	/*
			第一层 print()是 函数print的调用;第二层 func() string {} 和names... 分别是print的两个参数；
			结果
		0 |赵昌建|
		1 |kk|
		2 |17-赵|
	*/

	fmt.Println(strings.FieldsFunc("AbaabcAbcabc", func(ch rune) bool {
		return ch == 'a'
	}))
	//[Ab bcAbc bc]
	// 定义匿名函数并调用
	func() {
		fmt.Println("啦啦啦啦啦")
	}()
}

// 10 scan   输入的处理
// 原来讲到的  不带缓冲的 是说的语言级别

package main

import "fmt"

func main() {
	var num int
	fmt.Println(fmt.Scan(&num))
	fmt.Println(num)

}

// go for语句实现 死循环的语句
package main

import "fmt"

func main() {
	sum := 0
	idx := 1
	for {
		sum += idx
		idx++
		if idx > 100 { // 如果没有这个判断语句，上面就是一个死循环，通过这个条件，来判断退出
			break
		}
	}
	fmt.Println(sum, idx)
}

// go的循环语句中，没有while关键字，下面通过for if的组合 来实现while的效果，   在某某条件之前，循环会一直进行
package main

import "fmt"

func main() {
	sum := 0
	idx := 1
	for idx <= 100 { // 这个语句就有点类似别的语句中的  while  idx <= 100
		sum += idx
		idx++
	}
	fmt.Println(sum)
}

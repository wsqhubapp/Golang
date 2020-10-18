/*
需求
猜数字 生成随机数字0-100 从控制台数据 与生成数字比较 大 提示太大了 小 提示太小了 等于 成功, 程序结束
最多猜测五次，未猜对，说太笨了，程序结束
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number int
	number = rand.Intn(100)
	// fmt.Println(number)
	var gnumber int
	// idx := 1

	// if idx > 5 {
	// 	goto END
	// }
	// goto START
	// idx++

	// sum := 0
	// idx := 1
	// for idx <= 6 { // 这个语句就有点类似别的语句中的  while  idx <= 100
	// 	sum += idx
	// 	idx++
	// 	goto START
	// }
	// fmt.Println(sum)

	for idx := 1; idx < 4; idx++ {
		fmt.Println(idx)
		goto START

		// if idx > 5 {
		// 	fmt.Println(idx)
		// 	goto END
		// }
	}

START:
	fmt.Print("请输入一个数字:")
	fmt.Scan(&gnumber)
	fmt.Println("你输入的数字是：", gnumber)

	switch {
	case gnumber > number:
		fmt.Println("太大了")
		continue
	case gnumber < number:
		fmt.Println("太小了")
		continue
	default:
		fmt.Println("成功了")
	}

	// END:
	// 	fmt.Println("你太笨了")

}

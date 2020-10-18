// 逻辑判断  switch
/*实现逻辑判断
大于等于 90分的  得 A
大于等于 80分的  得 B
大于等于 70分的  得 C
大于等于 60分的  得 D
小于 60分的  得 E


*/
package main

import (
	"fmt"
)

func main() {
	fmt.Print("请输入你的分数: ")
	var score int
	fmt.Scan(&score)
	fmt.Printf("你输入的分数是 %d\n", score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}

/*
switch case  defalut 的基本写法， 条件判断写在case的行
*/

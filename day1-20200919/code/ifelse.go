package main

import "fmt"

func main() {
	//老公逻辑  如果有卖西瓜的 买一个包子，如果没有，就买十个包子

	var y string
	fmt.Print("有没有卖西瓜的： ")
	fmt.Scan(&y)
	fmt.Println("你的输入是： ", y)

	if y == "yes" {
		fmt.Println("买一个包子")
	} else {
		fmt.Println("买十个包子")
	}
}

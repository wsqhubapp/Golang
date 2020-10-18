package main

import "fmt"

func main() {
	var name string = "kk"
	fmt.Println("1", name)

	name = "silence" // 更新变量的值（赋值）
	fmt.Println("2", name)

	{
		// 不是定义，是赋值， 因为在这个块内，没有name变量，是找上一级，找到了变量name，就修改了，从而上一级（函数）的变量的值修改了
		name = "aaaaaaaaaaaaaaaaaa"
		fmt.Println("3", name)
	}
	fmt.Println("4", name)

}

/*
1 kk
2 silence
3 aaaaaaaaaaaaaaaaaa
4 aaaaaaaaaaaaaaaaaa  这个还是在函数中找值，变为aaaaaaaaaaa 是因为 name在块中赋值了（修改了）
*/

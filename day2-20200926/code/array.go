package main

import "fmt"

func main() {
	// 定义变量 names 元素类型为字符串长度为10 的数组
	// var names [10]string
	// var scores [10]int

	var names [5]string
	var scores [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%T, %#v\n", names, names)
	fmt.Printf("%T, %#v\n", scores, scores)
	fmt.Println(names, scores)
	// %q 显示值

	// 零值
	// n个对应元素类型的零值组成的数组
	// var scores [10]int

	// 索引
	// [...]type{i1:v1, ii:vi, in:vn}   长度是n+1
	// var scores [5]int = [5]int{1:8, 3:100}

	// 操作
	// 关系运算  !=  ==   相同的数据类型才能做关系运算？？！！

	// 访问值  修改值  基于索引来操作
<<<<<<< HEAD
	fmt.Println(names[0])
=======
	fmt.Println(nums[0])
>>>>>>> add

}

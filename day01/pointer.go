package main

import "fmt"

func main() {
	var (
		pointerInt    *int
		pointerString *string
	)

	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", pointerString, pointerString)

	// 赋值
	// 第一种方式 使用现有的变量取地址  方法是 &name
	// 这个复制的话，会在内存中重新开辟一段地址。 可以为理解为是深度复制
	age := 32

	age2 := age
	fmt.Printf("%T, %#v\n", &age, &age)
	fmt.Printf("%T, %#v\n", &age2, &age2)

	pointerInt = &age
	*pointerInt = 3300000
	fmt.Println(pointerInt)
	fmt.Printf("%T, %#v\n", &age2, &age2)

	//第二种方法  新值
	pointerString = new(string)
	fmt.Printf("%#v, %#v\n", pointerString, *pointerString)

	// 指针的指针 了解一下
	pp := &pointerString
	fmt.Printf("%T\n", pp)
	**pp = "kk"
	fmt.Println(*pointerString)

}

// 指针  在函数   结构体的内容用得比较多

// 指针的指针

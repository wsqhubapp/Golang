package main

import "fmt"

func main() {
	// isGirl := false

	// fmt.Printf("%T, %#v", isGirl, isGirl) // 结果是bool, false

	// 操作  主要是逻辑运算
	a, b, c, d := true, true, false, false
	// 与  && 左操作数与右操作数
	fmt.Println("a, b:", a && b) //
	fmt.Println("a, c:", a && c)
	fmt.Println("c, b:", c && b)
	fmt.Println("c, d:", c && d)

	// 或 ||
	fmt.Println("a, b:", a || b)
	fmt.Println("a, c:", a || c)
	fmt.Println("c, b:", c || b)
	fmt.Println("c, d:", c || d)

	// 非 ! 取反
	fmt.Println("a:", !a)
	fmt.Println("c:", !c)

	// 关系运算  == !=
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(c != b)
	fmt.Println(c == d)

	fmt.Printf("%t, %t\n", a, c) // 布尔占位符  %t

	var bbb bool          // 定义一个bool类型的零值
	fmt.Printf("%t", bbb) //  bool类型的零值是false
}

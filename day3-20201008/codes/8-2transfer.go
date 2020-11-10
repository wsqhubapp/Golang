/*
值传递
在 Go 语言中参数传递默认均为值传递（形参为实参变量的副本）
对于引用类型数据因其底层共享数据结构，
所以在函数内可对引用类型数据修改从而影响函数外的原变量信息

老师的说法   待定  ？？？？
	// 值传递(V)  go 中只有值传递  没有引用传递

	// 值类型 => 值拷贝
	// 引用类型 => 值拷贝(指针/地址) => 可能会影响函数外部的变量
*/
package main

import "fmt"

func main() {
	e1, e2 := "kk", []string{"kk", "silence"}
	// // 值传递
	// //在函数内修改值类型
	// fmt.Printf("e1: %p %v\n", &e1, e1)  //e1: 0xc00003c1f0 kk
	// func(e string) {
	// 	fmt.Printf("e: %p %v\n", &e, e)   //e: 0xc00003c210 kk
	// 	e = "silence"
	// }(e1)
	/*
		函数的调用过程，可以理解为赋值
		e := e1
		e = "silence"
	*/

	// //在函数内修改引用类型
	// fmt.Printf("e2: %p %v\n", &e2, e2)  //e2: 0xc0000044a0 [kk silence]
	// func(e []string) {
	// 	fmt.Printf("e: %p %v\n", &e, e) //e: 0xc000004520 [kk silence]
	// 	e[1] = "woniu"
	// }(e2)
	/*
		e : e2
		e[1] = "woniu"
	*/

	// fmt.Println(e1, e2)  //kk [kk woniu]

	//引用传递
	//在函数内修改值类型
	fmt.Printf("e1: %p %v\n", &e1, e1) //e1: 0xc00003c1f0 kk
	func(e *string) {
		fmt.Printf("e: %p %v\n", &e, *e) //e: 0xc000006030 kk
		*e = "silence"
	}(&e1)

	//在函数内修改引用类型
	fmt.Printf("e2: %p %v\n", &e2, e2) //e2: 0xc0000044a0 [kk silence]
	func(e *[]string) {
		fmt.Printf("e: %p %v\n", &e, *e) //e: 0xc000006038 [kk silence]
		(*e)[1] = "woniu"
	}(&e2)

	fmt.Println(e1, e2) //silence [kk woniu]
	/*
		引用传递
		可以通过将变量的地址通过指针类型传递给函数，此时可通过指针对函数外的原变量进行修改
	*/

}

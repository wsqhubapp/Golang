// 多个返回值和命名返回值

package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func add(a, b int) int {
	return a + b
}

// 返回多个值
// a + b, a-b , a* b, a/b
// go 语言支持函数有多个返回值，在声明函数时使用括号包含所有返回值类型，并使用return 返回对应数量的用逗号分隔数据
func op(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// 命名返回值
/*
在函数返回值列表中可指定变量名，变量在调用时会根据类型使用零值进行初始化，在函数中可进行赋值，同时在调用return时 不需要添加返回值
go语音自动将变量的最终结果返回
*/
func opv2(a, b int) (sum, sub, mul, div int) { // 类型同一个
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return //可以这么写，具体参考上面的说明
}

func main() {
	sayHello()
	add(1, 2)      // 没有输出  是因为虽然调用了函数，但是函数没有输出选项
	c := add(3, 4) // 7
	fmt.Println(c)
	fmt.Println(op(4, 2))
	a, b, c, d := op(4, 2) // 至少有一个变量未定义过  这里主要是说重新定了c
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d)

	fmt.Println(opv2(3, 2))
}

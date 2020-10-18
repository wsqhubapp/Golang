package main

import "fmt"

func main() {
	var age8 int8 = 31
	var age = 31
	fmt.Printf("%T, %#v, %d\n", age8, age8, age8)
	fmt.Printf("%T, %#v, %d\n", age, age, age)

	/*
		 了解
		8机制 0??  ?<8
		16进制 0x?? ?是0-9 a-f中的任意一个
		十进制的123
		1* 10^2 + 2 * 10^1 + 3* 10^0 = 100 + 20 + 3 =123  （进制转化的逻辑）

		2机制
		0111 转换为10机制
		1* 2^2 + 1*2^1 + 1*2^0 = 4+ 2 + 1 =7

		070 转换为10机制 ：这种0开头的为8进制，所以7*8^1 + 0 * 8^0 = 56 + 0 =56
		078 转换为10进制， 这种0开头的看着像8进制，实际看后面的8是不符合8进制中每一个数字小于8的要求，所以这个go程序中是一种不合规的写法，但是
		在实际的生活中，有10机制在前面加0 起到占位效果的。这个就是10机制 78
	*/
	//fmt.Println(070, 078)  // .\int.go:22:21: invalid digit '8' in octal literal

	// 常用操作
	// 算数运算 + - * / %  ++ --  加减乘除 取余 加加 减减；  要求数据类型一样； 除法运算要求是分母不能为零，遵从基本的数学知识。
	a, b := 2, 4
	fmt.Println(a / b) // 0
	fmt.Println(a % b) // 2
	a++
	b--

	fmt.Println(a) // 3
	fmt.Println(b) // 3

	// 关系运算 > < >= <= != ==
	fmt.Println(a > b) //fa
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a != b)
	fmt.Println(a == b)

	/* 了解  位运算
	负数 二进制表示补码，对应正数  取反+1
	-3 3=>0011  取反的结果是 1100 +1的结果是 1101


	按位与 &  两个都为1 结果为1 ， 7&2 = 2
	按位或 1  只要有一个为1 结果为1 ， 7|2 = 7
	取非: ^  2 0010 取非的话是1101  而 1101 表示-3
	右移位 >> 7 >> 2  0111 向右移动两位就是0001 结果就是1
	左移位 << 7<<2    0111 向左移动两位就是0000 0111 => 0001 1100 结果就是28
	and not &^ 7&^2 0111 0010 表示两个数的位数如果是相同的就取非 就变为了0101 就是5


	*/

	// 赋值运算 =  += -= /=  *=

	// 类型转换，必须是显示转换，不存在隐士转换，不转换就会报错  转换方法  type(value)  比如 int(int32)
	var (
		i   int   = 1
		i32 int32 = 1
		i64 int64 = 1
	)
	fmt.Printf("%T, %d\n", i+int(i32), i+int(i32))
	fmt.Printf("%T, %d\n", i+int(i32), i+int(i32))
	fmt.Printf("%T, %d\n", i32+int32(i), i32+int32(i))
	fmt.Printf("%T, %d\n", i64+int64(i), i64+int64(i))

	var (
		// 字符串存储到内存     转换涉及到编码的问题
		achar        byte = 'A'
		aint         byte = 65
		unicodePoint rune = '中'
	)

	fmt.Println(achar, aint)

	fmt.Printf("%d, %b, %o, %x, %U, %c, %c", achar, 15, 15, 15, unicodePoint, achar, 65)

}

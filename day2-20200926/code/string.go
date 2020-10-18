// 字符串
package main

import (
	"fmt"
	"utf8"
)

func main() {
	var name string = "kk"

	var desc string = "i love china"

	// 字面量
	// 可解析字符串  ""  原生字符串  `` 反引号

	// 零值
	// 操作
	// 连接
	// 关系运算符
	// 赋值操作
	// 长度  len  字节的长度  ascii
	// 索引   不能修改  索引的字节  ascii
	// 切片  得到的是字符串  对字节进行操作   ascii
	fmt.Println(name, desc)
	fmt.Println(desc[1:5])

	var text = "我爱中国"
	for i, v := range text {
		fmt.Printf("%d, %q\n", i, v)

	}

	fmt.Println(utf8.RnueCountInString(text))

}

package main

import "fmt"

func main() {
	s := "我爱中国"
	for i, v := range s {
		fmt.Printf("%d, %q\n", i, v)
	}

	for _, e := range s {
		fmt.Printf("%c\n", e)
	}

}

/*
for range  用于遍历可迭代对象中的每个元素，例如 字符串 数组 切片 映射 通道等
针对包含unicode字符的字符串遍历是需要使用for range range返回两个元素分别为字节索引和rune字符，可以通过空表标识符忽略要接收的变量
*/

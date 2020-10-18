package main

import "fmt"

func main() {
	var msg = "我的名字\\n是w"
	var msgRaw = `我的\n名字是w`
	fmt.Printf("%T %s\n", msg, msg)
	fmt.Printf("%T %s\n", msgRaw, msgRaw)

	// 操作
	// 字符串连接 +
	fmt.Println(msg + msgRaw)
	// 关系运算 > >= <=  != ==
	fmt.Println("abc" > "acd") //false

	// 赋值 = +=
	msg += "---w"
	fmt.Println(msg)

	// 切片 索引 ； 主要是针对ascii进行处理  得到的并不是第一个字符
	msg = "abcdef"
	fmt.Printf("%T %#v, %c\n", msg[0], msg[0], msg[0])
	fmt.Println(msg[1:3])

	// len  占用字节的大小
	fmt.Println(len(msg))    // 6
	fmt.Println(len(msgRaw)) // 18  #

}

// 函数嵌套
package main

import "fmt"

// 接收可变参数列表 string
// 函数格式化输出(针对可变参数列表中的每个元素) string->string
func print(formatter func(string) string, args ...string) {
	for i, v := range args {
		fmt.Println(i, formatter(v))
	}
}

// 第一层  func print() 这叫有参数，无返回值函数；第二层 参数中有两个；formatter func(string) string 和args ...string
// 第三层  函数有参数类型为string，返回值为string

func star(txt string) string {
	return "*" + txt + "*"
}

func table(txt string) string {
	return "|" + txt + "|"
}

func line(txt string, end string) string {
	return txt + end
}

func main() {
	names := []string{"赵昌建", "kk", "17-赵"}
	print(star, names...)
	// print(line, names...)
	print(table, names...)
}

/*
返回值的理解，把上面函数的三层理清楚  就好理解了
0 *赵昌建*
1 *kk*
2 *17-赵*

0 |赵昌建|
1 |kk|
2 |17-赵|
*/

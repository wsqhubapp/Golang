// 函数的类型
// 函数也可以赋值给变量，存储在数组，切片，映射中，也可以作为参数传递给函数或作为函数返回值返回

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func sayHi() {
	fmt.Println("hi")
}

func sayHello() {
	fmt.Println("hello")
}

func genFunc() func() { //返回值是一个函数类型 判断情况是奇数和偶数
	if rand.Int()%2 == 0 {
		return sayHi
	}
	return sayHello
}

func aFields(split rune) bool {
	return split == 'a'
}

func main() {
	rand.Seed(time.Now().Unix()) // 随机生成数据
	f := genFunc()
	f() // 函数调用

	fmt.Printf("%q\n", strings.FieldsFunc("abcdefabcabc", aFields)) // 调用字符串的分割函数

}

/*
10 main函数与init 函数

程序入口
开发者写代码的执行入口
*/

package main

import (
	"fmt"
	_ "testmain/pkg"
)

func main() {

}

// 初始化函数
// 导入包时执行
func init() {
	fmt.Println("main init")

}

/*
包的层级之间有顺序，包内没有顺序(程序不要把逻辑放这里)
*/

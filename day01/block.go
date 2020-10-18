package main

import "fmt"

// 包级别变量
var packageVar string = "package Var"

func main() {
	// 函数级别变量
	var funcVar string = "func Var"
	fmt.Println("1", packageVar, funcVar)

	{
		// 块级别变量
		var blockVar string = "block var"
		fmt.Println(packageVar, funcVar, blockVar)
		{
			// 子块级别变量
			var innerBlockVar string = "inner block var"
			fmt.Println(blockVar, innerBlockVar)
		}

	}
	fmt.Println("2", packageVar, funcVar)
	// fmt.Println(packageVar, funcVar, blockVar, innerBlockVar) //如果使用这个的话，blockVar, innerBlockVar 会没有定义的错误
}

/*
1 package Var func Var
package Var func Var block var
block var inner block var
2 package Var func Var

*/

/*
变量作用域
程序的执行顺序是从前到后的
子可以用父的（从下到上可以引用），父不能用子的变量
子可以重新定义一个变量，值会覆盖父里面的变量。也就是说值的寻找是从本区域（块）找，找到了，就不往上寻找；如果没有找到，则逐级往上寻找

*/

// 写法   大小写 ？

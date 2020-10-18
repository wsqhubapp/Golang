package testpkg

import "fmt"

// 包内成员(函数、变量、常量等)可见性
// 成员首字母大写包外可见
// 成员首字母小写只能包内可见

var T1Name = "T1"
var t3Name = "T3"

const T1Const = "T1Const"

func T1Func() {
	fmt.Println("T1func")
	t3Func()
}

func t3Func() {
	fmt.Println("t3func")
}

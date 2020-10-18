/*
为了演示 包名和目录名称不一致的问题
*/
package main

//导入目录

import (
	testpkg "8-testpkgname/pkg" //导入目录  //这是IDE帮忙写的别名
)

// 包名与文件名不一致
func main() {
	testpkg.Test() //包名调用
}

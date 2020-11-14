package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 开发终端  对CPU要求 有限制 比如 4核
	runtime.GOMAXPROCS(1) // 设置使用CPU核数
	fmt.Println(runtime.GOROOT())

}

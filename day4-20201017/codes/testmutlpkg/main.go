// 9 多个包之间调用的演示
package main

import (
	"testmutlpkg/pkga"
	"testmutlpkg/pkgb"
)

func main() {
	pkga.Test()
	pkgb.Test()
}

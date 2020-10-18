//12
package main

import (
	"flag"
	"fmt"
)

func main() {
	// ssh -P  port[22]  root
	var port int
	var help bool

	//指定变量与命令行参数(通过参数名称)中的关系
	flag.IntVar(&port, "P", 22, "port")
	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage: ssh -P 22 root@localhost")
		flag.PrintDefaults() //默认的方法
	}
	//解析命令行参数
	flag.Parse()

	if help {
		fmt.Println("帮助信息")
		fmt.Println("usage: ssh -P 22 root@loaclhoat")
		return
	}

	fmt.Println(port)
}

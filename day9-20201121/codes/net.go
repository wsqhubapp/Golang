package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("127.0.0.1", "8080")) //学习一个新的包的方法 先答应
	fmt.Println(net.SplitHostPort("127.0.0.1:888"))
	//ip -> 主机名
	fmt.Println(net.LookupAddr("127.0.0.1"))
	// 主机名 -》 ip
	fmt.Println(net.LookupHost("www.baidu.com"))
	fmt.Println(net.LookupHost("localhost"))
	fmt.Println(net.InterfaceAddrs())

}

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("127.0.0.1", "8080"))  //拼接主机名和端口 //学习一个新的包的方法 先打印
	fmt.Println(net.SplitHostPort("127.0.0.1:888")) 
	//ip -> 主机名
	fmt.Println(net.LookupAddr("127.0.0.1"))
	// 主机名 -》 ip
	fmt.Println(net.LookupHost("www.baidu.com"))
	fmt.Println(net.LookupHost("localhost"))
	fmt.Println(net.InterfaceAddrs())

	// IP, IPNet
	for _, ipStr := range []string{"127.0.0.1", "::1", "xxxx"} {
		ip := net.ParseIP(ipStr)
		fmt.Printf("%#v\n", ip)
	}

	ip, ipnet, err := net.ParseCIDR("192.168.0.0/24")
	// 192.168.0.0-192.168.0.255
	fmt.Printf("%#v %#v %#v\n", ip, ipnet, err)

	fmt.Println(ipnet.Contains(net.ParseIP("192.168.1.1")))  //false
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.0.10"))) //true

	addrs, err := net.InterfaceAddrs()
	fmt.Println(addrs, err)

	intfs, err := net.Interfaces()
	fmt.Println(intfs, err)
	for _, intf := range intfs {
		fmt.Println(intf.Index, intf.Name, intf.MTU, intf.HardwareAddr, intf.Flags)
		fmt.Println(intf.Addrs())
	}
	

}
/**/  net 包的一些基本用法



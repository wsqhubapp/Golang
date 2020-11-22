package main

import (
	"fmt"
	"net"
)

func main() {
	addr := "127.0.0.1:8888" //localhost

	protocol := "tcp"

	// 1 监听服务
	net.listener
	listener, err := net.listen(protocol, addr)
	fmt.Println(listener, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2 接受客户端请求
	conn, err := listener.Accept()
	// 3 与客户端交换数据
	conn.
	// 4 关闭客户端
	// 5 关闭服务
	net.interface

}

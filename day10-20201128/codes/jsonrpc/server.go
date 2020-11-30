package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 远程服务  Add(1, 2) int
type AddRequest struct {
	Left  int // 首字母大写是因为远程调用的可用  作用域的问题
	Right int
}

// type AddResponse int
type AddResponse struct {
	Result int
}

type Calc struct {
}

// 参数1 请求对象(可以是指针/值)
// 参数2 响应对象(可以是指针)
// 返回值：error
// func (c *Calc) Add(req AddRequest, resp *AddResponse) {
// 	resp.result = req.Left + req.Right
// }
// 错误处理  添加一个返回值
func (c *Calc) Add(req AddRequest, resp *AddResponse) error {
	fmt.Print("calc.add") //起到一个说明的作用
	resp.Result = req.Left + req.Right
	return nil
}

// type AAA struct {
// }

// func (a *AAA) Add(req AddRequest, resq *AddResponse) error {
// 	fmt.Println("aaa.Add")
// 	return nil
// }

func main() {
	rpc.Register(&Calc{})

	listener, _ := net.Listen("tcp", "0.0.0.0:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go jsonrpc.ServeConn(conn) // 开启例程

	}

	// listener.Close()
	// // 注册
	// rpc.Register(&Calc{})

	// // 绑定url和处理的关系
	// rpc.HandleHTTP()
	// http.ListenAndServe("0.0.0.0:8080", nil)

}

/*
问题扩展
比如 A包中有Calc的结构体   B 包中也有Calc的结构体
这样rpc.Register(&Calc{})  注册就不知道是调用哪个包的Calc 结构体
解决方法  rpc.Register("AAAA", &Calc{})  这样绑定就行
client.go  	err := client.Call("Calc.Add", req, &resp) 中的Calc.Add  修改为 AAAA.Add


*/

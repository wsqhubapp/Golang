package main

import (
	"fmt"
	"net"
	"net/rpc"
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

// //本地调用的时候的一个情况
// func main() {
// 	req := AddRequest{1, 2}  //结构体初始化
// 	resp := AddResponse{}    //结构体初始化
// 	calc := Calc{}           // 结构体初始化
// 	calc.Add(req, &resp)     // 方法的调用
// 	fmt.Println(resp.result) // 结果的输出  //3
// }

func main() {
	// rpc
	// 注册  （将上面的函数绑定到RPC的过程叫注册）
	rpc.Register(&Calc{}) //实现了基本的过程
	// rpc.Register(&AAA{})

	//启动服务  1) 建立监听  2）建立连接（默认实现了一个循环连接的功能）   3） 关闭
	listener, _ := net.Listen("tcp", "0.0.0.0:8080")
	rpc.Accept(listener)
	listener.Close()
}

/*
编码  解码

可以优化的地方： server.go  client.go 的两个结构体 都重复调用了；可以提到一个公共的文件中


*/

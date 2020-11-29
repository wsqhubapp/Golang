package main

import (
	"fmt"
	"net/http"
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

func main() {
	// 注册
	rpc.Register(&Calc{})

	// 绑定url和处理的关系
	rpc.HandleHTTP()
	http.ListenAndServe("0.0.0.0:8080", nil)

}

/*



 */

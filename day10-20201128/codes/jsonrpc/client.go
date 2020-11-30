package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"time"
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

func main() {
	// 建立连接
	client, _ := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	// client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//调用
	req := AddRequest{3, 100}
	resp := AddResponse{}
	err := client.Call("Calc.Add", req, &resp) // Call 用法参考文档   同步调用
	fmt.Println(err, resp)

	// 异步调用
	call := client.Go("Calc.Add", req, &resp, nil)
	for {
		select {
		case result := <-call.Done:
			fmt.Println(result.Reply, result.Error)
		default:
			fmt.Println(time.Now())
			time.Sleep(2 * time.Second)
		}
	}
	// 结束
	client.Close()
}

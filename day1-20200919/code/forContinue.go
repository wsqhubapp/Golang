package main

import "fmt"

func main() {
	// 输出
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 跳出本次循环，继续执行下一次循环
		}
		fmt.Println(i)

	}

}

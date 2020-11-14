//
package main

import (
	"fmt"
	"sync"
)

// 10个例程 分别给salary + 10  1000次
// 10个例程 分别给salary - 10  1000次

func main() {
	fmt.Println("start")
	var salary int = 0

	var locker sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			for i = 0; i < 1000; i++ {
				salary += 10
			}
		}()
		go func() {
			for i = 0; i < 1000; i++ {
				salary += 10
			}
		}()
	}

}

/*
如何保证数据不变---定义锁
锁保证在一次操作中数据是可控的。
信号计数器是控制+ - + -

sleep 无法预期时间
信号量和sleep 没什么区别

临时 用time.sleep  runtime  来控制
生产中用 wg   因为实际程序中无法预知时间  wg是一个状态的控制
*/

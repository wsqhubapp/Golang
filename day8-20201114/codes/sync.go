package main

import "sync"

func main() {

	/*
	   sync.WaitGroup  计数信号量
	   sync.Mutex  锁  互斥锁
	   sync.RWMutex  读写锁  这个锁比互斥锁性能高
	   sync.Cond  条件锁   多个历程   某个执行检查 是否满足条件  不满足等待
	   					其他历程 当可能产生等待历程条件重新满足 通知等待历程  Signal /Broadcase
	   sync.Map  # 线程安全的map  可以用map + sync.Mutex代替
	   sync.Once 保证一个逻辑只执行一次
	   sync.Pool 线程池  连接池   对象池  从池中获取对象，当池中无可用对象，创建并返回  当使用完成后会放入池中




	*/
	once := &sync.Once{}

	once.Do(task) // 函数只允许一次
	once.Do(taskA)

}

/*
new: string   结构体

*/

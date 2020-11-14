//

package main

func main() {
	fmt.Println("start")
	go func() {
		for i:='A'
	}
}

/*
go 后面加的是函数调用   

这里使用的是变量   没有使用指针？
没有赋值 就没有拷贝的过程  调用的是最初的wg 


信号量 没减少到零 go程序就会认为是锁住了

*/
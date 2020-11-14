package main

import "fmt"

//

func main() {
	// 声明int类型的管道
	// type 一般放值类型  不放引用类型
	// var channel chan int
	channel := make(chan int) // 无缓冲区管道  有人读才能写   

	// 初始化  必须使用make 初始化
	// make(chann int)

	//操作  读 写
	go func() {
		channel <- 1 //将1 写入管道

	}()



	num := <-channel //如果未读取到数据会进行阻塞  读管道
	fmt.Println(num)

	//关闭管道  关闭后不能写；关闭后能读取 但是读取的值是0
	close(channel)

	go func() {
		fmt.Println(<-channel)
	}()
	// 不能写
	// channel <- 1 、、会报错


	v, ok := <- channel
	fmt.Println(v, ok)


	//管道遍历
	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close（channel） 

	}()
	for num := range channel{  // 循环的话 如果上面没有停止，下面的循环就一直进行
		fmt.Println(num)
	}



}

/*
一个线程 在读的时候  没人写  会出现死锁  上面遍历的时候需要加 close 才解决这个问题
<- channel 读取数据 但是不管返回值



 */

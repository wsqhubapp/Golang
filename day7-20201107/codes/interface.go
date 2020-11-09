package main

import "fmt"

// 定义接口
/*
type 接口名称 interface {
	方法声明(行为) => 签名
}
*/

// 定义了一个Writer的接口
type Writer interface {
	Write([]byte) (int, error)
}

type MyFile struct { // 定义了一个MyFile的结构体
}

func (f *MyFile) write() { // 定义了一个方法
}

type MyFile2 struct {
}

func (f *MyFile2) Write(ctx []byte) (int, error) {
	return 0, nil
}

type MyFile3 struct {
}

func (f MyFile3) Write(ctx []byte) (int, error) {
	return 0, nil
}

// 自动生成指针接收者
/*
func (f *MyFile3) Write(ctx []byte) (int, error) {
	return 0, nil
}
*/

func main() {
	var writer Writer
	fmt.Printf("%T, %#v\n", writer, writer) //<nil>, <nil> 还没有赋值  这里就是空

	// var myFile MyFile
	// writer = myFile
	/*
			.\interface.go:28:9: cannot use myFile (type MyFile) as type Writer in assignment:
		        MyFile does not implement Writer (missing Write method)
		                have write()
						want Write([]byte) (int, error)
		这样的话 程序会报错，错误内容是myFile 没有Write 方法
	*/
	// var myFile2 MyFile2
	// writer = myFile2
	// 也会报错，因为write是值接收者  无Write 方法

	var myFile2Pointer = new(MyFile2)
	writer = myFile2Pointer
	fmt.Printf("%T, %#v\n", writer, writer) //*main.MyFile2, &main.MyFile2{} 指针类型

	var myFile3 MyFile3
	writer = myFile3
	fmt.Printf("%T, %#v\n", writer, writer)

	var myFile3Pointer = &MyFile3{} //new(MyFile3)
	writer = myFile3Pointer
	fmt.Printf("%T, %#v\n", writer, writer)

	// 结构体 MyFile3 变量赋值
	// 值类型变量 myFile3 := MyFile3{}
	//			 var myFile3 MyFile
	// 指针类型变量 var myFile3Pointer *MyFile3
	// myFile3Pointer = &myFile3
	// myFile3Pointer = &MyFile3{}
	// myFile3Pointer = new(MyFile3) new函数 创建结构体的指针类型变量

}

//创建值类型的时候自动创建指针类型的方法
//值类型自动实现了指针接收者方法

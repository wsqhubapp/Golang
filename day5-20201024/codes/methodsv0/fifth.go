//方法的值接收者和函数的值参数对比
package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

// 函数
func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

// 方法
func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}
func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)  //Area Function result: 50
	r.area() //Area Method result: 50

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p) //会报错

	p.area() //  Area Method result: 50  calling value receiver with a pointer
}

/*
当一个函数有一个值参数时，它只接受一个值参数。
当一个方法有一个值接收者时，它可以接受值和指针接收者
我们创建了一个指向 r 的指针 p。如果我们试图将这个指针传递给只接受值的 area 函数那么编译器将报错。
p.area() 使用指针接收者 p 调用一个值接收者方法 area 。这是完全合法的。
原因是对于 p.area()，由于 area 方法必须接受一个值接收者，所以 Golang 将其解析为 (*p).area()
*/

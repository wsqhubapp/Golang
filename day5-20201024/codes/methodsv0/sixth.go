//方法的指针接收者和函数的指针参数
package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

//函数  周长
func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

//方法
func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}
func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	p := &r       //pointer to r
	perimeter(p)  //perimeter function output: 30 #函数的参数 正常为指针类型
	p.perimeter() //perimeter method output: 30  #正常的指针接收者调用方法

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r)

	r.perimeter() //calling pointer receiver with a value //perimeter method output: 30 #Golang 会解析为(&r).perimeter()

}

/*
具有指针参数的函数将仅接受指针，而具有指针接收者的方法将接受值和指针接收者

试图以一个值参数 r 调用 perimeter 函数，这是非法的。因为一个接受指针为参数的函数不能接受一个值作为参数。如果去掉注释，则编译报错。
通过一个值接收者 r 调用一个指针接收者 perimeter 方法，这是合法的。r.perimeter() 将被 Golang 解析为 (&r).perimeter()。

*/

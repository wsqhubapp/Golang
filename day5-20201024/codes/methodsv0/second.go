//方法的加深理解
package main

import (
	"fmt"
	"math"
)

//定义长方形的结构体
type Rectangle struct {
	length int
	width  int
}

//定义圆形的结构体
type Circle struct {
	radius float64
}

//定义长方形计算面积的方法
func (r Rectangle) Area() int {
	return r.length * r.width
}

//定义圆计算面积的方法
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func main() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area()) //Area of rectangle 50
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area()) //Area of circle 452.389342
}

/*
长方形和圆形的结构体定义两个相同名字的方法，是可以的
*/

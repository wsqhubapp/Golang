//方法的定义  理解
package main

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}
func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type  //Salary of Sam Adolf is $5000
	//emp1 相当于实例化 一个实例调用一个方法
}

/*
也许有人会问，方法和函数差不多，为什么还要多此一举使用方法呢?
1 Golang 不是一个纯粹的面向对象的编程语言，它不支持类。因此通过在一个类型上建立方法来实现与 class 相似的行为。
2 同名方法可以定义在不同的类型上，但是 Golang 不允许同名函数。假设有两个结构体 Square 和 Circle。在 Square 和 Circle 上定义同名的方法是合法的。

方法的定义标准情况
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}
对各部分的说明：
接收器变量：接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母，而不是 self、this 之类的命名。例如，Socket 类型的接收器变量应该命名为 s，Connector 类型的接收器变量应该命名为 c 等。
接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：格式与函数定义一致。
*/

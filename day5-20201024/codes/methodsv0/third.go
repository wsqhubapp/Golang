// 方法的值接收者和指针接收者
// 两者区别在于，以指针作为接收者时，方法内部进行的修改对于调用者是可见的，但是以值作为接收者却不是
package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name) //Employee name before change: Mark Andrew
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name) //Employee name after change: Mark Andrew

	fmt.Printf("\n\nEmployee age before change: %d", e.age) //Employee age before change: 50
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age) //Employee age after change: 51
}

/*
值接收者  方法内部进行的修改对于调用者是不可见的，这是一个定义，原理的理解，可以再看看，这个就类似于赋值，相当于复制了
指针接收者  方法内部进行的修改对于调用者是可见的，这是一个定义，原理的理解，可以再看看，这个就类似于赋值，相当于复制了
使用 (&e).changeAge(51) 来调用 changeAge 方法不是必须的，Golang 允许我们省略 & 符号
因此可以写为 e.changeAge(51)。Golang 将 e.changeAge(51) 解析为 (&e).changeAge(51)。
*/
